// Copyright 2017 The Nxgit Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package misc

import (
	"go.khulnasoft.com/go-sdk/nxgit"
	"go.khulnasoft.com/nxgit/modules/context"
	"go.khulnasoft.com/nxgit/modules/setting"
)

// Version shows the version of the Nxgit server
func Version(ctx *context.APIContext) {
	// swagger:operation GET /version miscellaneous getVersion
	// ---
	// summary: Returns the version of the Nxgit application
	// produces:
	// - application/json
	// responses:
	//   "200":
	//     "$ref": "#/responses/ServerVersion"
	ctx.JSON(200, &nxgit.ServerVersion{Version: setting.AppVer})
}
