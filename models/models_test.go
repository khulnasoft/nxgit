// Copyright 2016 The Gogs Authors. All rights reserved.
// Copyright 2018 The Nxgit Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_parsePostgreSQLHostPort(t *testing.T) {
	tests := []struct {
		HostPort string
		Host     string
		Port     string
	}{
		{
			HostPort: "127.0.0.1:1234",
			Host:     "127.0.0.1",
			Port:     "1234",
		},
		{
			HostPort: "127.0.0.1",
			Host:     "127.0.0.1",
			Port:     "5432",
		},
		{
			HostPort: "[::1]:1234",
			Host:     "[::1]",
			Port:     "1234",
		},
		{
			HostPort: "[::1]",
			Host:     "[::1]",
			Port:     "5432",
		},
		{
			HostPort: "/tmp/pg.sock:1234",
			Host:     "/tmp/pg.sock",
			Port:     "1234",
		},
		{
			HostPort: "/tmp/pg.sock",
			Host:     "/tmp/pg.sock",
			Port:     "5432",
		},
	}
	for _, test := range tests {
		host, port := parsePostgreSQLHostPort(test.HostPort)
		assert.Equal(t, test.Host, host)
		assert.Equal(t, test.Port, port)
	}
}

func Test_getPostgreSQLConnectionString(t *testing.T) {
	tests := []struct {
		Host    string
		Port    string
		User    string
		Passwd  string
		Name    string
		Param   string
		SSLMode string
		Output  string
	}{
		{
			Host:    "/tmp/pg.sock",
			Port:    "4321",
			User:    "testuser",
			Passwd:  "space space !#$%^^%^```-=?=",
			Name:    "nxgit",
			Param:   "",
			SSLMode: "false",
			Output:  "postgres://testuser:space%20space%20%21%23$%25%5E%5E%25%5E%60%60%60-=%3F=@:5432/nxgitsslmode=false&host=/tmp/pg.sock",
		},
		{
			Host:    "localhost",
			Port:    "1234",
			User:    "pgsqlusername",
			Passwd:  "I love Nxgit!",
			Name:    "nxgit",
			Param:   "",
			SSLMode: "true",
			Output:  "postgres://pgsqlusername:I%20love%20Nxgit%21@localhost:5432/nxgitsslmode=true",
		},
	}

	for _, test := range tests {
		connStr := getPostgreSQLConnectionString(test.Host, test.User, test.Passwd, test.Name, test.Param, test.SSLMode)
		assert.Equal(t, test.Output, connStr)
	}
}
