---
date: "2017-04-08T11:34:00+02:00"
title: "Specific variables"
slug: "specific-variables"
weight: 20
toc: false
draft: false
menu:
  sidebar:
    parent: "advanced"
    name: "Specific variables"
    weight: 20
    identifier: "specific-variables"
---

# Specific variables

This is an inventory of Nxgit environment variables. They change Nxgit behaviour.

Initialize them before Nxgit command to be effective, for example:

```
NXGIT_CUSTOM=/home/nxgit/custom ./nxgit web
```

## From Go language

As Nxgit is written in Go, it uses some Go variables, such as:

  * `GOOS`
  * `GOARCH`
  * [`GOPATH`](https://golang.org/cmd/go/#hdr-GOPATH_environment_variable)

For documentation about each of the variables available, refer to the
[official Go documentation](https://golang.org/cmd/go/#hdr-Environment_variables).

## Nxgit files

  * `NXGIT_WORK_DIR`: Absolute path of working directory.
  * `NXGIT_CUSTOM`: Nxgit uses `NXGIT_WORK_DIR`/custom folder by default. Use this variable
     to change *custom* directory.
  * `GOGS_WORK_DIR`: Deprecated, use `NXGIT_WORK_DIR`
  * `GOGS_CUSTOM`: Deprecated, use `NXGIT_CUSTOM`

## Operating system specifics

  * `USER`: System user that Nxgit will run as. Used for some repository access strings.
  * `USERNAME`: if no `USER` found, Nxgit will use `USERNAME`
  * `HOME`: User home directory path. The `USERPROFILE` environment variable is used in Windows.

### Only on Windows

  * `USERPROFILE`: User home directory path. If empty, uses `HOMEDRIVE` + `HOMEPATH`
  * `HOMEDRIVE`: Main drive path used to access the home directory (C:)
  * `HOMEPATH`: Home relative path in the given home drive path

## Macaron (framework used by Nxgit)

  * `HOST`: Host Macaron will listen on
  * `PORT`: Port Macaron will listen on
  * `MACARON_ENV`: global variable to provide special functionality for development environments
     vs. production environments. If MACARON_ENV is set to "" or "development", then templates will
     be recompiled on every request. For more performance, set the MACARON_ENV environment variable
     to "production".

## Miscellaneous

  * `SKIP_MINWINSVC`: If set to 1, do not run as a service on Windows.
