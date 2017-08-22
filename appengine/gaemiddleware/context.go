// Copyright 2017 The LUCI Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package gaemiddleware

import (
	"net/http"

	"go.chromium.org/gae/filter/dscache"
	"go.chromium.org/gae/filter/readonly"
	"go.chromium.org/gae/service/datastore"

	"go.chromium.org/luci/common/data/caching/cacheContext"
	"go.chromium.org/luci/common/data/caching/lru"
	"go.chromium.org/luci/common/data/caching/proccache"
	"go.chromium.org/luci/common/logging"

	"go.chromium.org/luci/server/caching"
	"go.chromium.org/luci/server/middleware"
	"go.chromium.org/luci/server/router"
	"go.chromium.org/luci/server/settings"

	"go.chromium.org/luci/appengine/gaesecrets"
	"go.chromium.org/luci/appengine/gaesettings"

	"golang.org/x/net/context"
)

// Environment is a middleware environment. Its parameters define how the
// middleware is applied, and which services are enlisted.
type Environment struct {
	// PassthroughPanics, if true, instructs the Environment not to install panic
	// catching middleware.
	PassthroughPanics bool

	// DSDisableCache, if true, prevents the datastore caching filter from being
	// added.
	DSDisableCache bool

	// DSReadOnly, if true, causes a read-only datastore layer to be imposed,
	// preventing datastore writes.
	//
	// For any given datastore instance, at most one caching layer may be used.
	// All other instances must be ReadOnly to prevent errant writes from breaking
	// the assumptions of that caching layer. For example, if a Flex VM is being
	// used in conjunction with a non-read-only Classic AppEngine instance.
	DSReadOnly bool

	// WithInitialRequest is called at the very beginning of the handler. It
	// contains a reference to the handler's HTTP request.
	//
	// This should install basic services into the Context, including:
	// - Logging
	// - luci/GAE services.
	WithInitialRequest func(context.Context, *http.Request) context.Context

	// WithConfig is called during service setup to install the "gaeconfig" layer.
	//
	// If nil, no config layer will be installed.
	WithConfig func(context.Context) context.Context

	// WithAuth is called during service setup to install the "gaeauth" layer.
	//
	// If nil, no auth layer will be installed.
	WithAuth func(context.Context) context.Context

	// MonitoringMiddleware, if not nil, is additional middleware to append to the
	// end of the Base middleware chain to perform per-request monitoring.
	MonitoringMiddleware router.Middleware
}

var (
	// globalProcessCache holds state cached between requests.
	globalProcessCache = &proccache.Cache{}

	// globalLRUCache holds state cached between requests.
	//
	// TODO: We should choose a maximum LRU size here to stop this from growing
	//   indefinitely. However, it's replacing an indefinitely-growing cache, so
	//   for now this is acceptable.
	// TODO: This should replace uses of globalProcessCache.
	globalLRUCache = lru.New(0)

	// globalSettings holds global app settings lazily updated from the datastore.
	globalSettings = settings.New(gaesettings.Storage{})
)

// InstallHandlers installs handlers for an Environment's framework routes.
//
// See InstallHandlersWithMiddleware for more information.
func (e *Environment) InstallHandlers(r *router.Router) { InstallHandlersWithMiddleware(r, e.Base()) }

// With adds various production GAE LUCI services to the context.
//
// Basically, it installs GAE-specific backends and caches for various
// subsystems to make them work in GAE environment.
//
// One example is a backend for Logging: go.chromium.org/luci/common/logging.
// Logs emitted through a WithProd() context go to GAE logs.
//
// 'Production' here means the services will use real GAE APIs (not mocks or
// stubs), so With should never be used from unit tests.
func (e *Environment) With(c context.Context, req *http.Request) context.Context {
	// Set an initial logging level. We'll configure this to be more specific
	// later once we can load settings.
	c = logging.SetLevel(c, logging.Debug)
	c = caching.WithProcessCache(c, globalLRUCache)
	c = caching.WithRequestCache(c, nil)

	c = e.WithInitialRequest(c, req)

	// A previous layer must have installed a "luci/gae" datastore.
	if datastore.Raw(c) == nil {
		panic("no luci/gae datastore is installed")
	}

	// These are needed to use fetchCachedSettings.
	c = settings.Use(c, globalSettings)

	// Fetch and apply configuration stored in the datastore.
	cachedSettings := fetchCachedSettings(c)
	c = logging.SetLevel(c, cachedSettings.LoggingLevel)
	if !(e.DSDisableCache || bool(cachedSettings.DisableDSCache)) {
		c = dscache.AlwaysFilterRDS(c)
	}
	if e.DSReadOnly {
		c = readonly.FilterRDS(c)
	}

	// The rest of the service may use applied configuration.
	c = proccache.Use(c, globalProcessCache)
	if e.WithConfig != nil {
		c = e.WithConfig(c)
	}
	c = gaesecrets.Use(c, nil)
	if e.WithAuth != nil {
		c = e.WithAuth(c)
	}

	// Wrap this in a cache context so that lookups for any of the aforementioned
	// items are fast.
	return cacheContext.Wrap(c)
}

// Base returns a middleware chain to use for all GAE environment
// requests.
//
// Base DOES NOT install "luci/gae" services. To install appropriate services,
// use methods from a sub-package. This is done so that a given AppEngine
// environment doesn't need to include the superset of packages across all
// supported environments.
//
// This middleware chain installs prod GAE services into the request context
// (via With), and wraps the request with a panic catcher and monitoring
// hooks.
func (e *Environment) Base() router.MiddlewareChain {
	// addServices is a middleware that installs the set of standard production
	// AppEngine services by calling With.
	addServices := func(c *router.Context, next router.Handler) {
		// Create a cancelable Context that cancels when this request finishes.
		//
		// We do this because Contexts will leak resources and/or goroutines if they
		// have timers or can be canceled, but aren't. Predictably canceling the
		// parent will ensure that any such resource leaks are cleaned up.
		var cancelFunc context.CancelFunc
		c.Context, cancelFunc = context.WithCancel(c.Context)
		defer cancelFunc()

		// Apply production settings.
		c.Context = e.With(c.Context, c.Request)

		next(c)
	}

	mw := router.NewMiddlewareChain(addServices)
	if !e.PassthroughPanics {
		mw = mw.Extend(middleware.WithPanicCatcher)
	}
	if e.MonitoringMiddleware != nil {
		mw = mw.Extend(e.MonitoringMiddleware)
	}
	return mw
}
