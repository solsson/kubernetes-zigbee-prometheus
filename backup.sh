#!/bin/sh
set -e

[ -z "$BACKUP_PATH" ] && BACKUP_PATH=$(pwd)/backups
[ ! -d "$BACKUP_PATH" ] && echo "Backup path $BACKUP_PATH doesn't exist" && exit 1
[ -z "$NAME" ] && NAME=deconz-
[ -z "$LABEL" ] && LABEL=$(date -u +"%Y-%m-%d-%H%M%S")
[ -z "$DATA_PARENT" ] && DATA_PARENT=$(pwd)/testdata/dresden-elektronik
[ -z "$DATA_DIR" ] && DATA_DIR=deCONZ
[ ! -d "$DATA_PARENT/$DATA_DIR" ] && echo "Data dir $DATA_PARENT/$DATA_DIR not found" && exit 1

(cd $DATA_PARENT && tar cvfz $BACKUP_PATH/$NAME$LABEL.tgz $DATA_DIR)
