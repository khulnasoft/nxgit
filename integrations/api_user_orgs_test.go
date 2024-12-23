// Copyright 2018 The Nxgit Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.package models

package integrations

import (
	"fmt"
	"net/http"
	"testing"

	api "go.khulnasoft.com/go-sdk/nxgit"
	"go.khulnasoft.com/nxgit/models"

	"github.com/stretchr/testify/assert"
)

func TestUserOrgs(t *testing.T) {
	prepareTestEnv(t)
	adminUsername := "user1"
	normalUsername := "user2"
	session := loginUser(t, adminUsername)
	token := getTokenForLoggedInUser(t, session)
	urlStr := fmt.Sprintf("/api/v1/users/%s/orgs?token=%s", normalUsername, token)
	req := NewRequest(t, "GET", urlStr)
	resp := session.MakeRequest(t, req, http.StatusOK)
	var orgs []*api.Organization
	user3 := models.AssertExistsAndLoadBean(t, &models.User{Name: "user3"}).(*models.User)

	DecodeJSON(t, resp, &orgs)

	assert.Equal(t, []*api.Organization{
		{
			ID:          3,
			UserName:    user3.Name,
			FullName:    user3.FullName,
			AvatarURL:   user3.AvatarLink(),
			Description: "",
			Website:     "",
			Location:    "",
		},
	}, orgs)
}

func TestMyOrgs(t *testing.T) {
	prepareTestEnv(t)

	normalUsername := "user2"
	session := loginUser(t, normalUsername)
	token := getTokenForLoggedInUser(t, session)
	req := NewRequest(t, "GET", "/api/v1/user/orgs?token="+token)
	resp := session.MakeRequest(t, req, http.StatusOK)
	var orgs []*api.Organization
	DecodeJSON(t, resp, &orgs)
	user3 := models.AssertExistsAndLoadBean(t, &models.User{Name: "user3"}).(*models.User)

	assert.Equal(t, []*api.Organization{
		{
			ID:          3,
			UserName:    user3.Name,
			FullName:    user3.FullName,
			AvatarURL:   user3.AvatarLink(),
			Description: "",
			Website:     "",
			Location:    "",
		},
	}, orgs)
}
