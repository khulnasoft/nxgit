// Copyright 2017 The Nxgit Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package swagger

import (
	api "go.khulnasoft.com/go-sdk/nxgit"
)

// Organization
// swagger:response Organization
type swaggerResponseOrganization struct {
	// in:body
	Body api.Organization `json:"body"`
}

// OrganizationList
// swagger:response OrganizationList
type swaggerResponseOrganizationList struct {
	// in:body
	Body []api.Organization `json:"body"`
}

// Team
// swagger:response Team
type swaggerResponseTeam struct {
	// in:body
	Body api.Team `json:"body"`
}

// TeamList
// swagger:response TeamList
type swaggerResponseTeamList struct {
	// in:body
	Body []api.Team `json:"body"`
}
