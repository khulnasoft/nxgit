// Copyright 2014 The Gogs Authors. All rights reserved.
// Copyright 2016 The Nxgit Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"

	"go.khulnasoft.com/git"
	"go.khulnasoft.com/nxgit/models"
	"go.khulnasoft.com/nxgit/modules/log"
	"go.khulnasoft.com/nxgit/modules/pprof"
	"go.khulnasoft.com/nxgit/modules/private"
	"go.khulnasoft.com/nxgit/modules/setting"

	"github.com/Unknwon/com"
	"github.com/dgrijalva/jwt-go"
	version "github.com/mcuadros/go-version"
	"github.com/urfave/cli"
)

const (
	accessDenied        = "Repository does not exist or you do not have access"
	lfsAuthenticateVerb = "git-lfs-authenticate"
)

// CmdServ represents the available serv sub-command.
var CmdServ = cli.Command{
	Name:        "serv",
	Usage:       "This command should only be called by SSH shell",
	Description: `Serv provide access auth for repositories`,
	Action:      runServ,
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:  "config, c",
			Value: "custom/conf/app.ini",
			Usage: "Custom configuration file path",
		},
		cli.BoolFlag{
			Name: "enable-pprof",
		},
	},
}

func checkLFSVersion() {
	if setting.LFS.StartServer {
		//Disable LFS client hooks if installed for the current OS user
		//Needs at least git v2.1.2
		binVersion, err := git.BinVersion()
		if err != nil {
			fail(fmt.Sprintf("Error retrieving git version: %v", err), fmt.Sprintf("Error retrieving git version: %v", err))
		}

		if !version.Compare(binVersion, "2.1.2", ">=") {
			setting.LFS.StartServer = false
			println("LFS server support needs at least Git v2.1.2, disabled")
		} else {
			git.GlobalCommandArgs = append(git.GlobalCommandArgs, "-c", "filter.lfs.required=",
				"-c", "filter.lfs.smudge=", "-c", "filter.lfs.clean=")
		}
	}
}

func setup(logPath string) {
	log.DelLogger("console")
	setting.NewContext()
	checkLFSVersion()
	log.NewGitLogger(filepath.Join(setting.LogRootPath, logPath))
}

func parseCmd(cmd string) (string, string) {
	ss := strings.SplitN(cmd, " ", 2)
	if len(ss) != 2 {
		return "", ""
	}
	return ss[0], strings.Replace(ss[1], "'/", "'", 1)
}

var (
	allowedCommands = map[string]models.AccessMode{
		"git-upload-pack":    models.AccessModeRead,
		"git-upload-archive": models.AccessModeRead,
		"git-receive-pack":   models.AccessModeWrite,
		lfsAuthenticateVerb:  models.AccessModeNone,
	}
)

func fail(userMessage, logMessage string, args ...interface{}) {
	fmt.Fprintln(os.Stderr, "Nxgit:", userMessage)

	if len(logMessage) > 0 {
		if !setting.ProdMode {
			fmt.Fprintf(os.Stderr, logMessage+"\n", args...)
		}
		log.GitLogger.Fatal(3, logMessage, args...)
		return
	}

	log.GitLogger.Close()
	os.Exit(1)
}

func runServ(c *cli.Context) error {
	if c.IsSet("config") {
		setting.CustomConf = c.String("config")
	}
	setup("serv.log")

	if setting.SSH.Disabled {
		println("Nxgit: SSH has been disabled")
		return nil
	}

	if len(c.Args()) < 1 {
		cli.ShowSubcommandHelp(c)
		return nil
	}

	cmd := os.Getenv("SSH_ORIGINAL_COMMAND")
	if len(cmd) == 0 {
		println("Hi there, You've successfully authenticated, but Nxgit does not provide shell access.")
		println("If this is unexpected, please log in with password and setup Nxgit under another user.")
		return nil
	}

	verb, args := parseCmd(cmd)

	var lfsVerb string
	if verb == lfsAuthenticateVerb {
		if !setting.LFS.StartServer {
			fail("Unknown git command", "LFS authentication request over SSH denied, LFS support is disabled")
		}

		argsSplit := strings.Split(args, " ")
		if len(argsSplit) >= 2 {
			args = strings.TrimSpace(argsSplit[0])
			lfsVerb = strings.TrimSpace(argsSplit[1])
		}
	}

	repoPath := strings.ToLower(strings.Trim(args, "'"))
	rr := strings.SplitN(repoPath, "/", 2)
	if len(rr) != 2 {
		fail("Invalid repository path", "Invalid repository path: %v", args)
	}

	username := strings.ToLower(rr[0])
	reponame := strings.ToLower(strings.TrimSuffix(rr[1], ".git"))

	if setting.EnablePprof || c.Bool("enable-pprof") {
		if err := os.MkdirAll(setting.PprofDataPath, os.ModePerm); err != nil {
			fail("Error while trying to create PPROF_DATA_PATH", "Error while trying to create PPROF_DATA_PATH: %v", err)
		}

		stopCPUProfiler := pprof.DumpCPUProfileForUsername(setting.PprofDataPath, username)
		defer func() {
			stopCPUProfiler()
			pprof.DumpMemProfileForUsername(setting.PprofDataPath, username)
		}()
	}

	var (
		isWiki   bool
		unitType = models.UnitTypeCode
		unitName = "code"
	)
	if strings.HasSuffix(reponame, ".wiki") {
		isWiki = true
		unitType = models.UnitTypeWiki
		unitName = "wiki"
		reponame = reponame[:len(reponame)-5]
	}

	os.Setenv(models.EnvRepoUsername, username)
	if isWiki {
		os.Setenv(models.EnvRepoIsWiki, "true")
	} else {
		os.Setenv(models.EnvRepoIsWiki, "false")
	}
	os.Setenv(models.EnvRepoName, reponame)

	repo, err := private.GetRepositoryByOwnerAndName(username, reponame)
	if err != nil {
		if strings.Contains(err.Error(), "Failed to get repository: repository does not exist") {
			fail(accessDenied, "Repository does not exist: %s/%s", username, reponame)
		}
		fail("Internal error", "Failed to get repository: %v", err)
	}

	requestedMode, has := allowedCommands[verb]
	if !has {
		fail("Unknown git command", "Unknown git command %s", verb)
	}

	if verb == lfsAuthenticateVerb {
		if lfsVerb == "upload" {
			requestedMode = models.AccessModeWrite
		} else if lfsVerb == "download" {
			requestedMode = models.AccessModeRead
		} else {
			fail("Unknown LFS verb", "Unknown lfs verb %s", lfsVerb)
		}
	}

	// Prohibit push to mirror repositories.
	if requestedMode > models.AccessModeRead && repo.IsMirror {
		fail("mirror repository is read-only", "")
	}

	// Allow anonymous clone for public repositories.
	var (
		keyID int64
		user  *models.User
	)
	if requestedMode == models.AccessModeWrite || repo.IsPrivate || setting.Service.RequireSignInView {
		keys := strings.Split(c.Args()[0], "-")
		if len(keys) != 2 {
			fail("Key ID format error", "Invalid key argument: %s", c.Args()[0])
		}

		key, err := private.GetPublicKeyByID(com.StrTo(keys[1]).MustInt64())
		if err != nil {
			fail("Invalid key ID", "Invalid key ID[%s]: %v", c.Args()[0], err)
		}
		keyID = key.ID

		// Check deploy key or user key.
		if key.Type == models.KeyTypeDeploy {
			// Now we have to get the deploy key for this repo
			deployKey, err := private.GetDeployKey(key.ID, repo.ID)
			if err != nil {
				fail("Key access denied", "Failed to access internal api: [key_id: %d, repo_id: %d]", key.ID, repo.ID)
			}

			if deployKey == nil {
				fail("Key access denied", "Deploy key access denied: [key_id: %d, repo_id: %d]", key.ID, repo.ID)
			}

			if deployKey.Mode < requestedMode {
				fail("Key permission denied", "Cannot push with read-only deployment key: %d to repo_id: %d", key.ID, repo.ID)
			}

			// Update deploy key activity.
			if err = private.UpdateDeployKeyUpdated(key.ID, repo.ID); err != nil {
				fail("Internal error", "UpdateDeployKey: %v", err)
			}

			// FIXME: Deploy keys aren't really the owner of the repo pushing changes
			// however we don't have good way of representing deploy keys in hook.go
			// so for now use the owner
			os.Setenv(models.EnvPusherName, username)
			os.Setenv(models.EnvPusherID, fmt.Sprintf("%d", repo.OwnerID))
		} else {
			user, err = private.GetUserByKeyID(key.ID)
			if err != nil {
				fail("internal error", "Failed to get user by key ID(%d): %v", keyID, err)
			}

			if !user.IsActive || user.ProhibitLogin {
				fail("Your account is not active or has been disabled by Administrator",
					"User %s is disabled and have no access to repository %s",
					user.Name, repoPath)
			}

			mode, err := private.CheckUnitUser(user.ID, repo.ID, user.IsAdmin, unitType)
			if err != nil {
				fail("Internal error", "Failed to check access: %v", err)
			} else if *mode < requestedMode {
				clientMessage := accessDenied
				if *mode >= models.AccessModeRead {
					clientMessage = "You do not have sufficient authorization for this action"
				}
				fail(clientMessage,
					"User %s does not have level %v access to repository %s's "+unitName,
					user.Name, requestedMode, repoPath)
			}

			os.Setenv(models.EnvPusherName, user.Name)
			os.Setenv(models.EnvPusherID, fmt.Sprintf("%d", user.ID))
		}
	}

	//LFS token authentication
	if verb == lfsAuthenticateVerb {
		url := fmt.Sprintf("%s%s/%s.git/info/lfs", setting.AppURL, username, repo.Name)

		now := time.Now()
		claims := jwt.MapClaims{
			"repo": repo.ID,
			"op":   lfsVerb,
			"exp":  now.Add(setting.LFS.HTTPAuthExpiry).Unix(),
			"nbf":  now.Unix(),
		}
		if user != nil {
			claims["user"] = user.ID
		}
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

		// Sign and get the complete encoded token as a string using the secret
		tokenString, err := token.SignedString(setting.LFS.JWTSecretBytes)
		if err != nil {
			fail("Internal error", "Failed to sign JWT token: %v", err)
		}

		tokenAuthentication := &models.LFSTokenResponse{
			Header: make(map[string]string),
			Href:   url,
		}
		tokenAuthentication.Header["Authorization"] = fmt.Sprintf("Bearer %s", tokenString)

		enc := json.NewEncoder(os.Stdout)
		err = enc.Encode(tokenAuthentication)
		if err != nil {
			fail("Internal error", "Failed to encode LFS json response: %v", err)
		}

		return nil
	}

	// Special handle for Windows.
	if setting.IsWindows {
		verb = strings.Replace(verb, "-", " ", 1)
	}

	var gitcmd *exec.Cmd
	verbs := strings.Split(verb, " ")
	if len(verbs) == 2 {
		gitcmd = exec.Command(verbs[0], verbs[1], repoPath)
	} else {
		gitcmd = exec.Command(verb, repoPath)
	}
	if isWiki {
		if err = private.InitWiki(repo.ID); err != nil {
			fail("Internal error", "Failed to init wiki repo: %v", err)
		}
	}

	os.Setenv(models.ProtectedBranchRepoID, fmt.Sprintf("%d", repo.ID))

	gitcmd.Dir = setting.RepoRootPath
	gitcmd.Stdout = os.Stdout
	gitcmd.Stdin = os.Stdin
	gitcmd.Stderr = os.Stderr
	if err = gitcmd.Run(); err != nil {
		fail("Internal error", "Failed to execute git command: %v", err)
	}

	// Update user key activity.
	if keyID > 0 {
		if err = private.UpdatePublicKeyUpdated(keyID); err != nil {
			fail("Internal error", "UpdatePublicKey: %v", err)
		}
	}

	return nil
}
