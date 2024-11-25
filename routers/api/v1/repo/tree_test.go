// Copyright 2019 The Nxgit Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package repo

import (
	"github.com/stretchr/testify/assert"
	"testing"

	"go.khulnasoft.com/go-sdk/nxgit"
	"go.khulnasoft.com/nxgit/models"
	"go.khulnasoft.com/nxgit/modules/context"
	"go.khulnasoft.com/nxgit/modules/test"
)

func TestGetTreeBySHA(t *testing.T) {
	models.PrepareTestEnv(t)
	sha := "master"
	ctx := test.MockContext(t, "user2/repo1")
	ctx.SetParams(":id", "1")
	ctx.SetParams(":sha", sha)
	test.LoadRepo(t, ctx, 1)
	test.LoadRepoCommit(t, ctx)
	test.LoadUser(t, ctx, 2)
	test.LoadGitRepo(t, ctx)

	tree := GetTreeBySHA(&context.APIContext{Context: ctx, Org: nil}, ctx.Params("sha"))
	expectedTree := &nxgit.GitTreeResponse{
		SHA: "65f1bf27bc3bf70f64657658635e66094edbcb4d",
		URL: "https://try.nxgit.io/api/v1/repos/user2/repo1/git/trees/65f1bf27bc3bf70f64657658635e66094edbcb4d",
		Entries: []nxgit.GitEntry{
			{
				Path: "README.md",
				Mode: "100644",
				Type: "blob",
				Size: 30,
				SHA:  "4b4851ad51df6a7d9f25c979345979eaeb5b349f",
				URL:  "https://try.nxgit.io/api/v1/repos/user2/repo1/git/blobs/4b4851ad51df6a7d9f25c979345979eaeb5b349f",
			},
		},
		Truncated:  false,
		Page:       1,
		TotalCount: 1,
	}

	assert.EqualValues(t, tree, expectedTree)
}
