//go:build !bindata
// +build !bindata

// Copyright 2016 The Nxgit Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package public

import (
	"gopkg.in/macaron.v1"
)

// Static implements the macaron static handler for serving assets.
func Static(opts *Options) macaron.Handler {
	return opts.staticHandler(opts.Directory)
}
