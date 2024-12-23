// Copyright 2019 The Nxgit Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package mail

import (
	"go.khulnasoft.com/nxgit/models"
	"go.khulnasoft.com/nxgit/modules/log"
	"go.khulnasoft.com/nxgit/modules/notification/base"
)

type mailNotifier struct {
	base.NullNotifier
}

var (
	_ base.Notifier = &mailNotifier{}
)

// NewNotifier create a new mailNotifier notifier
func NewNotifier() base.Notifier {
	return &mailNotifier{}
}

func (m *mailNotifier) NotifyCreateIssueComment(doer *models.User, repo *models.Repository,
	issue *models.Issue, comment *models.Comment) {
	var act models.ActionType
	if comment.Type == models.CommentTypeClose {
		act = models.ActionCloseIssue
	} else if comment.Type == models.CommentTypeReopen {
		act = models.ActionReopenIssue
	} else if comment.Type == models.CommentTypeComment {
		act = models.ActionCommentIssue
	} else if comment.Type == models.CommentTypeCode {
		act = models.ActionCommentIssue
	}

	if err := comment.MailParticipants(act, issue); err != nil {
		log.Error(4, "MailParticipants: %v", err)
	}
}

func (m *mailNotifier) NotifyNewIssue(issue *models.Issue) {
	if err := issue.MailParticipants(); err != nil {
		log.Error(4, "MailParticipants: %v", err)
	}
}

func (m *mailNotifier) NotifyIssueChangeStatus(doer *models.User, issue *models.Issue, isClosed bool) {
	if err := issue.MailParticipants(); err != nil {
		log.Error(4, "MailParticipants: %v", err)
	}
}

func (m *mailNotifier) NotifyNewPullRequest(pr *models.PullRequest) {
	if err := pr.Issue.MailParticipants(); err != nil {
		log.Error(4, "MailParticipants: %v", err)
	}
}

func (m *mailNotifier) NotifyPullRequestReview(pr *models.PullRequest, r *models.Review, comment *models.Comment) {
	var act models.ActionType
	if comment.Type == models.CommentTypeClose {
		act = models.ActionCloseIssue
	} else if comment.Type == models.CommentTypeReopen {
		act = models.ActionReopenIssue
	} else if comment.Type == models.CommentTypeComment {
		act = models.ActionCommentIssue
	}
	if err := comment.MailParticipants(act, pr.Issue); err != nil {
		log.Error(4, "MailParticipants: %v", err)
	}
}
