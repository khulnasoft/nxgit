---
date: "2018-05-07T13:00:00+02:00"
title: "Nxgit compared to other Git hosting options"
slug: "comparison"
weight: 5
toc: true
draft: false
menu:
  sidebar:
    parent: "features"
    name: "Comparison"
    weight: 5
    identifier: "comparison"
---

# Nxgit compared to other Git hosting options

To help decide if Nxgit is suited for your needs, here is how it compares to other Git self hosted options.

Be warned that we don't regularly check for feature changes in other products, so this list may be outdated. If you find anything that needs to be updated in the table below, please report it in an [issue on GitHub](https://github.com/khulnasoft/nxgit/issues).

_Symbols used in table:_

* _✓ - supported_

* _⁄ - supported with limited functionality_

* _✘ - unsupported_

#### General Features

| Feature | Nxgit | Gogs | GitHub EE | GitLab CE | GitLab EE | BitBucket | RhodeCode CE |
|---------|-------|------|-----------|-----------|-----------|-----------|--------------|
| Open source and free | ✓ | ✓ | ✘| ✓ | ✘ | ✘ | ✓ |
| Low resource usage (RAM/CPU) | ✓ | ✓ | ✘ | ✘ | ✘ | ✘ | ✘ |
| Multiple database support | ✓ | ✓ | ✘ | ⁄ | ⁄ | ✓ | ✓ |
| Multiple OS support | ✓ | ✓ | ✘ | ✘ | ✘ | ✘ | ✓ |
| Easy upgrade process | ✓ | ✓ | ✘ | ✓ | ✓ | ✘ | ✓ |
| Markdown support | ✓ | ✓ | ✓ | ✓ | ✓ | ✓ | ✓ |
| Orgmode support | ✓ | ✘ | ✓ | ✘ | ✘ | ✘ | ? |
| CSV support | ✓ | ✘ | ✓ | ✘ | ✘ | ✓ | ? |
| Third-party render tool support | ✓ | ✘ | ✘ | ✘ | ✘ | ✓ | ? |
| Static Git-powered pages | ✘ | ✘ | ✓ | ✓ | ✓ | ✘ | ✘ |
| Integrated Git-powered wiki | ✓ | ✓ | ✓ | ✓ | ✓ | ✓ | ✘ |
| Deploy Tokens | ✓ | ✓ | ✓ | ✓ | ✓ | ✓ | ✓ |
| Repository Tokens with write rights | ✓ | ✘ | ✓ | ✓ | ✓ | ✘ | ✓ |
| Built-in Container Registry | ✘ | ✘ | ✘ | ✓ | ✓ | ✘ | ✘ |
| External git mirroring | ✓ | ✓ | ✘ | ✘ | ✓ | ✓ | ✓ |
| FIDO U2F (2FA) | ✓ | ✘ | ✓ | ✓ | ✓ | ✓ | ✘ |
| Built-in CI/CD | ✘ | ✘ | ✘ | ✓ | ✓ | ✘ | ✘ |
| Subgroups: groups within groups | ✘ | ✘ | ✘ | ✓ | ✓ | ✘ | ✓ |

#### Code management

| Feature | Nxgit | Gogs | GitHub EE | GitLab CE | GitLab EE | BitBucket | RhodeCode CE |
|---------|-------|------|-----------|-----------|-----------|-----------|--------------|
| Repository topics | ✓ | ✘ | ✓ | ✓ | ✓ | ✘ | ✘ |
| Repository code search | ✓ | ✘ | ✓ | ✓ | ✓ | ✓ | ✓ |
| Global code search | ✓ | ✘ | ✓ | ✘ | ✓ | ✓ | ✓ |
| Git LFS 2.0 | ✓ | ✘ | ✓ | ✓ | ✓ | ⁄ | ✓ |
| Group Milestones | ✘ | ✘ | ✘ | ✓ | ✓ | ✘ | ✘ |
| Granular user roles (Code, Issues, Wiki etc) | ✓ | ✘ | ✘ | ✓ | ✓ | ✘ | ✘ |
| Verified Committer | ✘ | ✘ | ? | ✓ | ✓ | ✓ | ✘ |
| GPG Signed Commits | ✓ | ✘ | ✓ | ✓ | ✓ | ✓ | ✓ |
| Reject unsigned commits | ✘ | ✘ | ✓ | ✓ | ✓ | ✘ | ✓ |
| Repository Activity page | ✓ | ✘ | ✓ | ✓ | ✓ | ✓ | ✓ |
| Branch manager | ✓ | ✘ | ✓ | ✓ | ✓ | ✓ | ✓ |
| Create new branches | ✓ | ✘ | ✓ | ✓ | ✓ | ✘ | ✘ |
| Web code editor | ✓ | ✓ | ✓ | ✓ | ✓ | ✓ | ✓ |
| Commit graph | ✓ | ✘ | ✓ | ✓ | ✓ | ✓ | ✓ |

#### Issue Tracker

| Feature | Nxgit | Gogs | GitHub EE | GitLab CE | GitLab EE | BitBucket | RhodeCode CE |
|---------|-------|------|-----------|-----------|-----------|-----------|--------------|
| Issue tracker | ✓ | ✓ | ✓ | ✓ | ✓ | ✓ | ✘ |
| Issue templates | ✓ | ✓ | ✓ | ✓ | ✓ | ✘ | ✘ |
| Labels | ✓ | ✓ | ✓ | ✓ | ✓ | ✘ | ✘ |
| Time tracking | ✓ | ✘ | ✓ | ✓ | ✓ | ✘ | ✘ |
| Multiple assignees for issues | ✓ | ✘ | ✓ | ✘ | ✓ | ✘ | ✘ |
| Related issues | ✘ | ✘ | ⁄ | ✘ | ✓ | ✘ | ✘ |
| Confidential issues | ✘ | ✘ | ✘ | ✓ | ✓ | ✘ | ✘ |
| Comment reactions | ✓ | ✘ | ✓ | ✓ | ✓ | ✘ | ✘ |
| Lock Discussion | ✓ | ✘ | ✓ | ✓ | ✓ | ✘ | ✘ |
| Batch issue handling | ✓ | ✘ | ✓ | ✓ | ✓ | ✘ | ✘ |
| Issue Boards | ✘ | ✘ | ✘ | ✓ | ✓ | ✘ | ✘ |
| Create new branches from issues | ✘ | ✘ | ✘ | ✓ | ✓ | ✘ | ✘ |
| Issue search | ✓ | ✘ | ✓ | ✓ | ✓ | ✓ | ✘ |
| Global issue search | ✘ | ✘ | ✓ | ✓ | ✓ | ✓ | ✘ |
| Issue dependency | ✓ | ✘ | ✘ | ✘ | ✘ | ✘ | ✘ |
| Create issue via email | [✘](https://github.com/khulnasoft/nxgit/issues/6226) | [✘](https://github.com/gogs/gogs/issues/2602) | ✘ | ✘ | ✓ | ✓ | ✘ |
| Service Desk | [✘](https://github.com/khulnasoft/nxgit/issues/6219) | ✘ | ✘ | ✘ | ✓ | ✘ | ✘ |

#### Pull/Merge requests

| Feature | Nxgit | Gogs | GitHub EE | GitLab CE | GitLab EE | BitBucket | RhodeCode CE |
|---------|-------|------|-----------|-----------|-----------|-----------|--------------|
| Pull/Merge requests | ✓ | ✓ | ✓ | ✓ | ✓ | ✓ | ✓ |
| Squash merging | ✓ | ✘ | ✓ | ✘ | ✓ | ✓ | ✓ |
| Rebase merging | ✓ | ✓ | ✓ | ✘ | ⁄ | ✘ | ✓ |
| Pull/Merge request inline comments | ✓ | ✘ | ✓ | ✓ | ✓ | ✓ | ✓ |
| Pull/Merge request approval | ✓ | ✘ | ⁄ | ✓ | ✓ | ✓ | ✓ |
| Merge conflict resolution | ✘ | ✘ | ✓ | ✓ | ✓ | ✓ | ✘ |
| Restrict push and merge access to certain users | ✓ | ✘ | ✓ | ⁄ | ✓ | ✓ | ✓ |
| Revert specific commits or a merge request | ✘ | ✘ | ✓ | ✓ | ✓ | ✓ | ✘ |
| Pull/Merge requests templates | ✓ | ✓ | ✓ | ✓ | ✓ | ✘ | ✘ |
| Cherry-picking changes | ✘ | ✘ | ✘ | ✓ | ✓ | ✘ | ✘ |


#### 3rd-party integrations

| Feature | Nxgit | Gogs | GitHub EE | GitLab CE | GitLab EE | BitBucket | RhodeCode CE |
|---------|-------|------|-----------|-----------|-----------|-----------|--------------|
| Webhook support | ✓ | ✓ | ✓ | ✓ | ✓ | ✓ | ✓ |
| Custom Git Hooks | ✓ | ✓ | ✓ | ✓ | ✓ | ✓ | ✓ |
| AD / LDAP integration | ✓ | ✓ | ✓ | ✓ | ✓ | ✓ | ✓ |
| Multiple LDAP / AD server support | ✓ | ✓ | ✘ | ✘ | ✓ | ✓ | ✓ |
| LDAP user synchronization | ✓ | ✘ | ✓ | ✓ | ✓ | ✓ | ✓ |
| OpenId Connect support | ✓ | ✘ | ✓ | ✓ | ✓ | ? | ✘ |
| OAuth 2.0 integration (external authorization) | ✓ | ✘ | ⁄ | ✓ | ✓ | ? | ✓ |
| Act as OAuth 2.0 provider | [✓](https://github.com/khulnasoft/nxgit/pull/5378) | ✘ | ✓ | ✓ | ✓ | ✓ | ✘ |
| Two factor authentication (2FA) | ✓ | ✓ | ✓ | ✓ | ✓ | ✓ | ✘ |
| Mattermost/Slack integration | ✓ | ✓ | ⁄ | ✓ | ✓ | ⁄ | ✓ |
| Discord integration | ✓ | ✓ | ✓ | ✓ | ✓ | ✘ | ✘ |
| External CI/CD status display | ✓ | ✘ | ✓ | ✓ | ✓ | ✓ | ✓ |
