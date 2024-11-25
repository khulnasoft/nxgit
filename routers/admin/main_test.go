// Copyright 2018 The Nxgit Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package admin

import (
	"path/filepath"
	"testing"

	"go.khulnasoft.com/nxgit/models"
)

func TestMain(m *testing.M) {
	models.MainTest(m, filepath.Join("..", ".."))
}
