// Copyright 2017 The Nxgit Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package integrations

import (
	"net/http"
	"testing"
)

func TestExploreRepos(t *testing.T) {
	prepareTestEnv(t)

	req := NewRequest(t, "GET", "/explore/repos")
	MakeRequest(t, req, http.StatusOK)
}
