// Copyright 2016 Nxgit. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package migrations

import (
	"fmt"

	"github.com/go-xorm/xorm"
)

func createAllowCreateOrganizationColumn(x *xorm.Engine) error {
	type User struct {
		KeepEmailPrivate        bool
		AllowCreateOrganization bool
	}

	if err := x.Sync2(new(User)); err != nil {
		return fmt.Errorf("Sync2: %v", err)
	} else if _, err = x.Where("`type` = 0").Cols("allow_create_organization").Update(&User{AllowCreateOrganization: true}); err != nil {
		return fmt.Errorf("set allow_create_organization: %v", err)
	}
	return nil
}
