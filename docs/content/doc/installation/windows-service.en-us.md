---
date: "2016-12-21T15:00:00-02:00"
title: "Register as a Windows Service"
slug: "windows-service"
weight: 10
toc: true
draft: false
menu:
  sidebar:
    parent: "installation"
    name: "Windows Service"
    weight: 30
    identifier: "windows-service"
---

# Prerequisites

The following changes are made in C:\nxgit\custom\conf\app.ini:

```
RUN_USER = COMPUTERNAME$
```

Sets Nxgit to run as the local system user.

COMPUTERNAME is whatever the response is from `echo %COMPUTERNAME%` on the command line. If the response is `USER-PC` then `RUN_USER = USER-PC$`

## Use absolute paths

If you use sqlite3, change the `PATH` to include the full path:

```
[database]
PATH     = c:/nxgit/data/nxgit.db
```

# Register as a Windows Service

To register Nxgit as a Windows service, open a command prompt (cmd) as an Administrator,
then run the following command:

```
sc create nxgit start= auto binPath= ""C:\nxgit\nxgit.exe" web --config "C:\nxgit\custom\conf\app.ini""
```

Do not forget to replace `C:\nxgit` with the correct Nxgit directory.

Open "Windows Services", search for the service named "nxgit", right-click it and click on
"Run". If everything is OK, Nxgit will be reachable on `http://localhost:3000` (or the port
that was configured).

## Unregister as a service

To unregister Nxgit as a service, open a command prompt (cmd) as an Administrator and run:

```
sc delete nxgit
```
