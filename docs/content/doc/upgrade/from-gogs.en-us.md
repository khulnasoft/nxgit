---
date: "2016-12-01T16:00:00+02:00"
title: "Upgrade from Gogs"
slug: "upgrade-from-gogs"
weight: 10
toc: true
draft: false
menu:
  sidebar:
    parent: "upgrade"
    name: "From Gogs"
    weight: 10
    identifier: "upgrade-from-gogs"
---

# Upgrade from Gogs

Gogs, version 0.9.146 and older, can be easily migrated to Nxgit.

There are some basic steps to follow. On a Linux system run as the Gogs user:

* Create a Gogs backup with `gogs backup`. This creates `gogs-backup-[timestamp].zip` file
  containing all important Gogs data. You would need it if you wanted to move to the `gogs` back later.
* Download the file matching the destination platform from the [downloads page](https://dl.nxgit.io/nxgit).
 It should be `1.0.x` version. Migrating from `gogs` to any other version is impossible.
* Put the binary at the desired install location.
* Copy `gogs/custom/conf/app.ini` to `nxgit/custom/conf/app.ini`.
* Copy custom `templates, public` from `gogs/custom/` to `nxgit/custom/`.
* For any other custom folders, such as `gitignore, label, license, locale, readme` in
  `gogs/custom/conf`, copy them to `nxgit/custom/options`.
* Copy `gogs/data/` to `nxgit/data/`. It contains issue attachments and avatars.
* Verify by starting Nxgit with `nxgit web`.
* Enter Nxgit admin panel on the UI, run `Rewrite '.ssh/authorized_keys' file`.
* Launch every major version of the binary ( `1.1.4` → `1.2.3` → `1.3.4` → `1.4.2` →  etc ) to migrate database.
* If custom or config path was changed, run `Rewrite all update hook of repositories`.

## Change gogs specific information

* Rename `gogs-repositories/` to `nxgit-repositories/`
* Rename `gogs-data/` to `nxgit-data/`
* In `nxgit/custom/conf/app.ini` change:

  FROM:

  ```ini
  [database]
  PATH = /home/:USER/gogs/data/:DATABASE.db
  [attachment]
  PATH = /home/:USER/gogs-data/attachments
  [picture]
  AVATAR_UPLOAD_PATH = /home/:USER/gogs-data/avatars
  [log]
  ROOT_PATH = /home/:USER/gogs/log
  ```

  TO:

  ```ini
  [database]
  PATH = /home/:USER/nxgit/data/:DATABASE.db
  [attachment]
  PATH = /home/:USER/nxgit-data/attachments
  [picture]
  AVATAR_UPLOAD_PATH = /home/:USER/nxgit-data/avatars
  [log]
  ROOT_PATH = /home/:USER/nxgit/log
  ```

* Verify by starting Nxgit with `nxgit web`

## Upgrading to most recent `nxgit` version

After successful migration from `gogs` to `nxgit 1.0.x`, it is possible to upgrade to the recent `nxgit` version.
Simply download the file matching the destination platform from the [downloads page](https://dl.nxgit.io/nxgit)
and replace the binary.

## Troubleshooting

* If errors are encountered relating to custom templates in the `nxgit/custom/templates`
  folder, try moving the templates causing the errors away one by one. They may not be
  compatible with Nxgit or an update.

## Add Nxgit to startup on Unix

Update the appropriate file from [nxgit/contrib](https://github.com/khulnasoft/nxgit/tree/master/contrib)
with the right environment variables.

For distros with systemd:

* Copy the updated script to `/etc/systemd/system/nxgit.service`
* Add the service to the startup with: `sudo systemctl enable nxgit`
* Disable old gogs startup script: `sudo systemctl disable gogs`

For distros with SysVinit:

* Copy the updated script to `/etc/init.d/nxgit`
* Add the service to the startup with: `sudo rc-update add nxgit`
* Disable old gogs startup script: `sudo rc-update del gogs`
