// Copyright 2017 The Nxgit Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package validation

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/go-macaron/binding"
)

const (
	// ErrGitRefName is git reference name error
	ErrGitRefName = "GitRefNameError"
)

var (
	// GitRefNamePattern is regular expression with unallowed characters in git reference name
	GitRefNamePattern = regexp.MustCompile("[^\\d\\w-_\\./]")
)

// AddBindingRules adds additional binding rules
func AddBindingRules() {
	addGitRefNameBindingRule()
	addValidURLBindingRule()
}

func addGitRefNameBindingRule() {
	// Git refname validation rule
	binding.AddRule(&binding.Rule{
		IsMatch: func(rule string) bool {
			return strings.HasPrefix(rule, "GitRefName")
		},
		IsValid: func(errs binding.Errors, name string, val interface{}) (bool, binding.Errors) {
			str := fmt.Sprintf("%v", val)

			if GitRefNamePattern.MatchString(str) {
				errs.Add([]string{name}, ErrGitRefName, "GitRefName")
				return false, errs
			}
			// Additional rules as described at https://www.kernel.org/pub/software/scm/git/docs/git-check-ref-format.html
			if strings.HasPrefix(str, "/") || strings.HasSuffix(str, "/") ||
				strings.HasSuffix(str, ".") || strings.Contains(str, "..") ||
				strings.Contains(str, "//") {
				errs.Add([]string{name}, ErrGitRefName, "GitRefName")
				return false, errs
			}
			parts := strings.Split(str, "/")
			for _, part := range parts {
				if strings.HasSuffix(part, ".lock") || strings.HasPrefix(part, ".") {
					errs.Add([]string{name}, ErrGitRefName, "GitRefName")
					return false, errs
				}
			}

			return true, errs
		},
	})
}

func addValidURLBindingRule() {
	// URL validation rule
	binding.AddRule(&binding.Rule{
		IsMatch: func(rule string) bool {
			return strings.HasPrefix(rule, "ValidUrl")
		},
		IsValid: func(errs binding.Errors, name string, val interface{}) (bool, binding.Errors) {
			str := fmt.Sprintf("%v", val)
			if len(str) != 0 && !IsValidURL(str) {
				errs.Add([]string{name}, binding.ERR_URL, "Url")
				return false, errs
			}

			return true, errs
		},
	})
}

func portOnly(hostport string) string {
	colon := strings.IndexByte(hostport, ':')
	if colon == -1 {
		return ""
	}
	if i := strings.Index(hostport, "]:"); i != -1 {
		return hostport[i+len("]:"):]
	}
	if strings.Contains(hostport, "]") {
		return ""
	}
	return hostport[colon+len(":"):]
}

func validPort(p string) bool {
	for _, r := range []byte(p) {
		if r < '0' || r > '9' {
			return false
		}
	}
	return true
}
