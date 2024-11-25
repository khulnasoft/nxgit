// Copyright 2017 The Nxgit Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package integrations

import (
	"net/http"
	"testing"

	"go.khulnasoft.com/go-sdk/nxgit"
	"go.khulnasoft.com/nxgit/modules/setting"

	"github.com/stretchr/testify/assert"
)

func TestVersion(t *testing.T) {
	prepareTestEnv(t)

	setting.AppVer = "test-version-1"
	req := NewRequest(t, "GET", "/api/v1/version")
	resp := MakeRequest(t, req, http.StatusOK)

	var version nxgit.ServerVersion
	DecodeJSON(t, resp, &version)
	assert.Equal(t, setting.AppVer, string(version.Version))
}
