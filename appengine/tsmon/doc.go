// Copyright 2016 The LUCI Authors. All rights reserved.
// Use of this source code is governed under the Apache License, Version 2.0
// that can be found in the LICENSE file.

// Package tsmon adapts common/tsmon library to GAE environment.
//
// It configures tsmon state with a monitor and store suitable for GAE
// environment and controls when metric flushes happen.
package tsmon