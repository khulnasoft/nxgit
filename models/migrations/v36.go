// Copyright 2017 The Nxgit Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package migrations

import (
	"go.khulnasoft.com/nxgit/models"

	"github.com/go-xorm/xorm"
)

func regenerateGitHooks36(x *xorm.Engine) (err error) {
	return models.SyncRepositoryHooks()
}
