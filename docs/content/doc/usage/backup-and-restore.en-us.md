---
date: "2017-01-01T16:00:00+02:00"
title: "Usage: Backup and Restore"
slug: "backup-and-restore"
weight: 11
toc: true
draft: false
menu:
  sidebar:
    parent: "usage"
    name: "Backup and Restore"
    weight: 11
    identifier: "backup-and-restore"
---

# Backup and Restore

Nxgit currently has a `dump` command that will save the installation to a zip file. This
file can be unpacked and used to restore an instance.

## Backup Command (`dump`)

Switch to the user running nxgit: `su git`. Run `./nxgit dump -c /path/to/app.ini` in the nxgit installation
directory. There should be some output similar to the following:

```
2016/12/27 22:32:09 Creating tmp work dir: /tmp/nxgit-dump-417443001
2016/12/27 22:32:09 Dumping local repositories.../home/git/nxgit-repositories
2016/12/27 22:32:22 Dumping database...
2016/12/27 22:32:22 Packing dump files...
2016/12/27 22:32:34 Removing tmp work dir: /tmp/nxgit-dump-417443001
2016/12/27 22:32:34 Finish dumping in file nxgit-dump-1482906742.zip
```

Inside the `nxgit-dump-1482906742.zip` file, will be the following:

* `custom` - All config or customerize files in `custom/`.
* `data` - Data directory in <NXGIT_WORK_DIR>, except sessions if you are using file session. This directory includes `attachments`, `avatars`, `lfs`, `indexers`, sqlite file if you are using sqlite.
* `nxgit-db.sql` - SQL dump of database
* `nxgit-repo.zip` - Complete copy of the repository directory.
* `log/` - Various logs. They are not needed for a recovery or migration.

Intermediate backup files are created in a temporary directory specified either with the
`--tempdir` command-line parameter or the `TMPDIR` environment variable.

## Restore Command (`restore`)

There is currently no support for a recovery command. It is a manual process that mostly
involves moving files to their correct locations and restoring a database dump.

Example:
```
apt-get install nxgit
unzip nxgit-dump-1482906742.zip
cd nxgit-dump-1482906742
mv custom/conf/app.ini /etc/nxgit/conf/app.ini
unzip nxgit-repo.zip
mv nxgit-repo/* /var/lib/nxgit/repositories/
chown -R nxgit:nxgit /etc/nxgit/conf/app.ini /var/lib/nxgit/repositories/
mysql -u$USER -p$PASS $DATABASE <nxgit-db.sql
# or  sqlite3 $DATABASE_PATH <nxgit-db.sql
service nxgit restart
```
