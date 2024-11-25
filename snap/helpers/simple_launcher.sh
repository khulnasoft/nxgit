#!/bin/bash

if ! env | grep -q root; then
  echo "
   +----------------------------------------+
   | You are not running nxgit as root.     |
   | This is required for the snap package. |
   | Please re-run as root.                 |
   +----------------------------------------+
"
  $SNAP/nxgit/nxgit --help
  exit 1
fi

# Set usernames for nxgit
export USERNAME=root
export USER=root

export NXGIT_WORK_DIR=$(snapctl get nxgit.snap.workdir)
export NXGIT_CUSTOM=$(snapctl get nxgit.snap.custom)

$SNAP/bin/gconfig save
cd $SNAP/nxgit; ./nxgit $@
