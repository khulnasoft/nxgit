// Copyright 2017 Nxgit. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package models

import (
	"container/list"
	"fmt"
	"strings"

	"go.khulnasoft.com/git"
	api "go.khulnasoft.com/go-sdk/nxgit"
	"go.khulnasoft.com/nxgit/modules/log"
	"go.khulnasoft.com/nxgit/modules/setting"
	"go.khulnasoft.com/nxgit/modules/util"

	"github.com/go-xorm/xorm"
)

// CommitStatusState holds the state of a Status
// It can be "pending", "success", "error", "failure", and "warning"
type CommitStatusState string

// IsWorseThan returns true if this State is worse than the given State
func (css CommitStatusState) IsWorseThan(css2 CommitStatusState) bool {
	switch css {
	case CommitStatusError:
		return true
	case CommitStatusFailure:
		return css2 != CommitStatusError
	case CommitStatusWarning:
		return css2 != CommitStatusError && css2 != CommitStatusFailure
	case CommitStatusSuccess:
		return css2 != CommitStatusError && css2 != CommitStatusFailure && css2 != CommitStatusWarning
	default:
		return css2 != CommitStatusError && css2 != CommitStatusFailure && css2 != CommitStatusWarning && css2 != CommitStatusSuccess
	}
}

const (
	// CommitStatusPending is for when the Status is Pending
	CommitStatusPending CommitStatusState = "pending"
	// CommitStatusSuccess is for when the Status is Success
	CommitStatusSuccess CommitStatusState = "success"
	// CommitStatusError is for when the Status is Error
	CommitStatusError CommitStatusState = "error"
	// CommitStatusFailure is for when the Status is Failure
	CommitStatusFailure CommitStatusState = "failure"
	// CommitStatusWarning is for when the Status is Warning
	CommitStatusWarning CommitStatusState = "warning"
)

// CommitStatus holds a single Status of a single Commit
type CommitStatus struct {
	ID          int64             `xorm:"pk autoincr"`
	Index       int64             `xorm:"INDEX UNIQUE(repo_sha_index)"`
	RepoID      int64             `xorm:"INDEX UNIQUE(repo_sha_index)"`
	Repo        *Repository       `xorm:"-"`
	State       CommitStatusState `xorm:"VARCHAR(7) NOT NULL"`
	SHA         string            `xorm:"VARCHAR(64) NOT NULL INDEX UNIQUE(repo_sha_index)"`
	TargetURL   string            `xorm:"TEXT"`
	Description string            `xorm:"TEXT"`
	Context     string            `xorm:"TEXT"`
	Creator     *User             `xorm:"-"`
	CreatorID   int64

	CreatedUnix util.TimeStamp `xorm:"INDEX created"`
	UpdatedUnix util.TimeStamp `xorm:"INDEX updated"`
}

func (status *CommitStatus) loadRepo(e Engine) (err error) {
	if status.Repo == nil {
		status.Repo, err = getRepositoryByID(e, status.RepoID)
		if err != nil {
			return fmt.Errorf("getRepositoryByID [%d]: %v", status.RepoID, err)
		}
	}
	if status.Creator == nil && status.CreatorID > 0 {
		status.Creator, err = getUserByID(e, status.CreatorID)
		if err != nil {
			return fmt.Errorf("getUserByID [%d]: %v", status.CreatorID, err)
		}
	}
	return nil
}

// APIURL returns the absolute APIURL to this commit-status.
func (status *CommitStatus) APIURL() string {
	status.loadRepo(x)
	return fmt.Sprintf("%sapi/v1/%s/statuses/%s",
		setting.AppURL, status.Repo.FullName(), status.SHA)
}

// APIFormat assumes some fields assigned with values:
// Required - Repo, Creator
func (status *CommitStatus) APIFormat() *api.Status {
	status.loadRepo(x)
	apiStatus := &api.Status{
		Created:     status.CreatedUnix.AsTime(),
		Updated:     status.CreatedUnix.AsTime(),
		State:       api.StatusState(status.State),
		TargetURL:   status.TargetURL,
		Description: status.Description,
		ID:          status.Index,
		URL:         status.APIURL(),
		Context:     status.Context,
	}
	if status.Creator != nil {
		apiStatus.Creator = status.Creator.APIFormat()
	}

	return apiStatus
}

// CalcCommitStatus returns commit status state via some status, the commit statues should order by id desc
func CalcCommitStatus(statuses []*CommitStatus) *CommitStatus {
	var lastStatus *CommitStatus
	var state CommitStatusState
	for _, status := range statuses {
		if status.State.IsWorseThan(state) {
			state = status.State
			lastStatus = status
		}
	}
	if lastStatus == nil {
		if len(statuses) > 0 {
			lastStatus = statuses[0]
		} else {
			lastStatus = &CommitStatus{}
		}
	}
	return lastStatus
}

// GetCommitStatuses returns all statuses for a given commit.
func GetCommitStatuses(repo *Repository, sha string, page int) ([]*CommitStatus, error) {
	statuses := make([]*CommitStatus, 0, 10)
	return statuses, x.Limit(10, page*10).Where("repo_id = ?", repo.ID).And("sha = ?", sha).Find(&statuses)
}

// GetLatestCommitStatus returns all statuses with a unique context for a given commit.
func GetLatestCommitStatus(repo *Repository, sha string, page int) ([]*CommitStatus, error) {
	ids := make([]int64, 0, 10)
	err := x.Limit(10, page*10).
		Table(&CommitStatus{}).
		Where("repo_id = ?", repo.ID).And("sha = ?", sha).
		Select("max( id ) as id").
		GroupBy("context").OrderBy("max( id ) desc").Find(&ids)
	if err != nil {
		return nil, err
	}
	statuses := make([]*CommitStatus, 0, len(ids))
	if len(ids) == 0 {
		return statuses, nil
	}
	return statuses, x.In("id", ids).Find(&statuses)
}

// GetCommitStatus populates a given status for a given commit.
// NOTE: If ID or Index isn't given, and only Context, TargetURL and/or Description
//
//	is given, the CommitStatus created _last_ will be returned.
func GetCommitStatus(repo *Repository, sha string, status *CommitStatus) (*CommitStatus, error) {
	conds := &CommitStatus{
		Context:     status.Context,
		State:       status.State,
		TargetURL:   status.TargetURL,
		Description: status.Description,
	}
	has, err := x.Where("repo_id = ?", repo.ID).And("sha = ?", sha).Desc("created_unix").Get(conds)
	if err != nil {
		return nil, fmt.Errorf("GetCommitStatus[%s, %s]: %v", repo.RepoPath(), sha, err)
	}
	if !has {
		return nil, fmt.Errorf("GetCommitStatus[%s, %s]: not found", repo.RepoPath(), sha)
	}

	return conds, nil
}

// NewCommitStatusOptions holds options for creating a CommitStatus
type NewCommitStatusOptions struct {
	Repo         *Repository
	Creator      *User
	SHA          string
	CommitStatus *CommitStatus
}

func newCommitStatus(sess *xorm.Session, opts NewCommitStatusOptions) error {
	opts.CommitStatus.Description = strings.TrimSpace(opts.CommitStatus.Description)
	opts.CommitStatus.Context = strings.TrimSpace(opts.CommitStatus.Context)
	opts.CommitStatus.TargetURL = strings.TrimSpace(opts.CommitStatus.TargetURL)
	opts.CommitStatus.SHA = opts.SHA
	opts.CommitStatus.CreatorID = opts.Creator.ID

	if opts.Repo == nil {
		return fmt.Errorf("newCommitStatus[nil, %s]: no repository specified", opts.SHA)
	}
	opts.CommitStatus.RepoID = opts.Repo.ID
	repoPath := opts.Repo.repoPath(sess)

	if opts.Creator == nil {
		return fmt.Errorf("newCommitStatus[%s, %s]: no user specified", repoPath, opts.SHA)
	}

	gitRepo, err := git.OpenRepository(repoPath)
	if err != nil {
		return fmt.Errorf("OpenRepository[%s]: %v", repoPath, err)
	}
	if _, err := gitRepo.GetCommit(opts.SHA); err != nil {
		return fmt.Errorf("GetCommit[%s]: %v", opts.SHA, err)
	}

	// Get the next Status Index
	var nextIndex int64
	lastCommitStatus := &CommitStatus{
		SHA:    opts.SHA,
		RepoID: opts.Repo.ID,
	}
	has, err := sess.Desc("index").Limit(1).Get(lastCommitStatus)
	if err != nil {
		sess.Rollback()
		return fmt.Errorf("newCommitStatus[%s, %s]: %v", repoPath, opts.SHA, err)
	}
	if has {
		log.Debug("newCommitStatus[%s, %s]: found", repoPath, opts.SHA)
		nextIndex = lastCommitStatus.Index
	}
	opts.CommitStatus.Index = nextIndex + 1
	log.Debug("newCommitStatus[%s, %s]: %d", repoPath, opts.SHA, opts.CommitStatus.Index)

	// Insert new CommitStatus
	if _, err = sess.Insert(opts.CommitStatus); err != nil {
		sess.Rollback()
		return fmt.Errorf("newCommitStatus[%s, %s]: %v", repoPath, opts.SHA, err)
	}

	return nil
}

// NewCommitStatus creates a new CommitStatus given a bunch of parameters
// NOTE: All text-values will be trimmed from whitespaces.
// Requires: Repo, Creator, SHA
func NewCommitStatus(repo *Repository, creator *User, sha string, status *CommitStatus) error {
	sess := x.NewSession()
	defer sess.Close()

	if err := sess.Begin(); err != nil {
		return fmt.Errorf("NewCommitStatus[repo_id: %d, user_id: %d, sha: %s]: %v", repo.ID, creator.ID, sha, err)
	}

	if err := newCommitStatus(sess, NewCommitStatusOptions{
		Repo:         repo,
		Creator:      creator,
		SHA:          sha,
		CommitStatus: status,
	}); err != nil {
		return fmt.Errorf("NewCommitStatus[repo_id: %d, user_id: %d, sha: %s]: %v", repo.ID, creator.ID, sha, err)
	}

	return sess.Commit()
}

// SignCommitWithStatuses represents a commit with validation of signature and status state.
type SignCommitWithStatuses struct {
	Status *CommitStatus
	*SignCommit
}

// ParseCommitsWithStatus checks commits latest statuses and calculates its worst status state
func ParseCommitsWithStatus(oldCommits *list.List, repo *Repository) *list.List {
	var (
		newCommits = list.New()
		e          = oldCommits.Front()
	)

	for e != nil {
		c := e.Value.(SignCommit)
		commit := SignCommitWithStatuses{
			SignCommit: &c,
		}
		statuses, err := GetLatestCommitStatus(repo, commit.ID.String(), 0)
		if err != nil {
			log.Error(3, "GetLatestCommitStatus: %v", err)
		} else {
			commit.Status = CalcCommitStatus(statuses)
		}

		newCommits.PushBack(commit)
		e = e.Next()
	}
	return newCommits
}
