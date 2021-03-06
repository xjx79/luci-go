// Copyright 2017 The LUCI Authors. All rights reserved.
// Use of this source code is governed under the Apache License, Version 2.0
// that can be found in the LICENSE file.

package frontend

import (
	"go.chromium.org/luci/milo/buildsource"
	"go.chromium.org/luci/server/router"
	"go.chromium.org/luci/server/templates"
)

// LogHandler is responsible for taking a universal build ID and rendering the
// build page (defined in ./appengine/templates/pages/log.html).
func LogHandler(c *router.Context, buildID buildsource.ID, logname string) error {
	log, closed, err := buildID.GetLog(c.Context, logname)
	if err != nil {
		return err
	}
	templates.MustRender(c.Context, c.Writer, "pages/log.html", templates.Args{
		"Log":    log,
		"Closed": closed,
	})
	return nil
}
