// Copyright 2014 The Gogs Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package admin

import (
	"go.khulnasoft.com/nxgit/models"
	"go.khulnasoft.com/nxgit/modules/base"
	"go.khulnasoft.com/nxgit/modules/context"
	"go.khulnasoft.com/nxgit/modules/setting"
	"go.khulnasoft.com/nxgit/routers"
)

const (
	tplOrgs base.TplName = "admin/org/list"
)

// Organizations show all the organizations
func Organizations(ctx *context.Context) {
	ctx.Data["Title"] = ctx.Tr("admin.organizations")
	ctx.Data["PageIsAdmin"] = true
	ctx.Data["PageIsAdminOrganizations"] = true

	routers.RenderUserSearch(ctx, &models.SearchUserOptions{
		Type:     models.UserTypeOrganization,
		PageSize: setting.UI.Admin.OrgPagingNum,
		Private:  true,
	}, tplOrgs)
}
