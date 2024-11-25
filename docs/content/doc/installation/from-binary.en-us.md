---
date: "2017-06-19T12:00:00+02:00"
title: "Installation from binary"
slug: "install-from-binary"
weight: 10
toc: true
draft: false
menu:
  sidebar:
    parent: "installation"
    name: "From binary"
    weight: 20
    identifier: "install-from-binary"
---

# Installation from binary

All downloads come with SQLite, MySQL and PostgreSQL support, and are built with
embedded assets. This can be different for older releases. Choose the file matching
the destination platform from the [downloads page](https://dl.nxgit.io/nxgit/), copy
the URL and replace the URL within the commands below:

```sh
wget -O nxgit https://dl.nxgit.io/nxgit/1.7.0/nxgit-1.7.0-linux-amd64
chmod +x nxgit
```

## Verify GPG signature
Nxgit signs all binaries with a [GPG key](https://pgp.mit.edu/pks/lookup?op=vindex&fingerprint=on&search=0x2D9AE806EC1592E2) to prevent against unwanted modification of binaries. To validate the binary, download the signature file which ends in `.asc` for the binary you downloaded and use the gpg command line tool.

```sh
gpg --keyserver pgp.mit.edu --recv 7C9E68152594688862D62AF62D9AE806EC1592E2
gpg --verify nxgit-1.7.0-linux-amd64.asc nxgit-1.7.0-linux-amd64
```

## Test

After getting a binary, it can be tested with `./nxgit web` or moved to a permanent
location. When launched manually, Nxgit can be killed using `Ctrl+C`.

```
./nxgit web
```

## Recommended server configuration

**NOTE:** Many of the following directories can be configured using [Environment Variables]({{< relref "doc/advanced/specific-variables.en-us.md" >}}) as well!  
Of note, configuring `NXGIT_WORK_DIR` will tell Nxgit where to base its working directory, as well as ease installation.

### Prepare environment

Check that Git is installed on the server. If it is not, install it first.
```sh
git --version
```

Create user to run Nxgit (ex. `git`)
```sh
adduser \
   --system \
   --shell /bin/bash \
   --gecos 'Git Version Control' \
   --group \
   --disabled-password \
   --home /home/git \
   git
```

### Create required directory structure

```sh
mkdir -p /var/lib/nxgit/{custom,data,log}
chown -R git:git /var/lib/nxgit/
chmod -R 750 /var/lib/nxgit/
mkdir /etc/nxgit
chown root:git /etc/nxgit
chmod 770 /etc/nxgit
```

**NOTE:** `/etc/nxgit` is temporary set with write rights for user `git` so that Web installer could write configuration file. After installation is done, it is recommended to set rights to read-only using:
```
chmod 750 /etc/nxgit
chmod 644 /etc/nxgit/app.ini
```

### Configure Nxgit's working directory

**NOTE:** If you plan on running Nxgit as a Linux service, you can skip this step as the service file allows you to set `WorkingDirectory`. Otherwise, consider setting this environment variable (semi-)permanently so that Nxgit consistently uses the correct working directory.
```
export NXGIT_WORK_DIR=/var/lib/nxgit/
```

### Copy Nxgit binary to global location

```
cp nxgit /usr/local/bin/nxgit
```

## Running Nxgit

After the above steps, two options to run Nxgit are:

### 1. Creating a service file to start Nxgit automatically (recommended)

See how to create [Linux service]({{< relref "run-as-service-in-ubuntu.en-us.md" >}})

### 2. Running from command-line/terminal

```
NXGIT_WORK_DIR=/var/lib/nxgit/ /usr/local/bin/nxgit web -c /etc/nxgit/app.ini
```

## Updating to a new version

You can update to a new version of Nxgit by stopping Nxgit, replacing the binary at `/usr/local/bin/nxgit` and restarting the instance. 
The binary file name should not be changed during the update to avoid problems 
in existing repositories. 

It is recommended you do a [backup]({{< relref "doc/usage/backup-and-restore.en-us.md" >}}) before updating your installation.

If you have carried out the installation steps as described above, the binary should 
have the generic name `nxgit`. Do not change this, i.e. to include the version number. 

See below for troubleshooting instructions to repair broken repositories after 
an update of your Nxgit version.

## Troubleshooting

### Old glibc versions

Older Linux distributions (such as Debian 7 and CentOS 6) may not be able to load the
Nxgit binary, usually producing an error such as ```./nxgit: /lib/x86_64-linux-gnu/libc.so.6:
version `GLIBC\_2.14' not found (required by ./nxgit)```. This is due to the integrated
SQLite support in the binaries provided by dl.nxgit.io. In this situation, it is usually
possible to [install from source]({{< relref "from-source.en-us.md" >}}) without sqlite
support.

### Running Nxgit on another port

For errors like `702 runWeb()] [E] Failed to start server: listen tcp 0.0.0.0:3000:
bind: address already in use` Nxgit needs to be started on another free port. This
is possible using `./nxgit web -p $PORT`. It's possible another instance of Nxgit
is already running.

### Git error after updating to a new version of Nxgit

If the binary file name has been changed during the update to a new version of Nxgit, 
git hooks in existing repositories will not work any more. In that case, a git 
error will be displayed when pushing to the repository.

```
remote: ./hooks/pre-receive.d/nxgit: line 2: [...]: No such file or directory
```

The `[...]` part of the error message will contain the path to your previous Nxgit 
binary.

To solve this, go to the admin options and run the task `Resynchronize pre-receive, 
update and post-receive hooks of all repositories` to update all hooks to contain
the new binary path. Please note that this overwrite all git hooks including ones
with customizations made.

If you aren't using the built-in to Nxgit SSH server you will also need to re-write
the authorized key file by running the `Update the '.ssh/authorized_keys' file with
Nxgit SSH keys.` task in the admin options.
