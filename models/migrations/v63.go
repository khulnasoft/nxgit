// Copyright 2018 The Nxgit Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package migrations

import (
	"fmt"

	"github.com/go-xorm/xorm"
)

func addLanguageSetting(x *xorm.Engine) error {
	type User struct {
		Language string `xorm:"VARCHAR(5)"`
	}

	if err := x.Sync2(new(User)); err != nil {
		return fmt.Errorf("Sync2: %v", err)
	}

	return nil
}
