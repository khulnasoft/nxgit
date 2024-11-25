// Copyright 2017 The Nxgit Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package misc

import (
	"go.khulnasoft.com/nxgit/modules/base"
	"go.khulnasoft.com/nxgit/modules/context"
)

// tplSwagger swagger page template
const tplSwagger base.TplName = "swagger/ui"

// Swagger render swagger-ui page with v1 json
func Swagger(ctx *context.Context) {
	ctx.Data["APIJSONVersion"] = "v1"
	ctx.HTML(200, tplSwagger)
}
