#!/bin/bash

# not go alone if crash
set -e

DBUSER=root
DBHOST=95.163.202.160
DBNAME=homework
DBPASSWORD=vimi

#get remote database sql to local
function dump() {
	mysqldump -h$DBHOST -u$DBUSER -p$DBPASSWORD $DBNAME > ./db/sql/latest_dump.sql
}

function restore() {
	mysql -h$DBHOST -u$DBUSER -p$DBPASSWORD $DBNAME < ./db/sql/latest_dump.sql
}

case "$1" in
	dump)
		dump
		;;

	restore)
		restore
		;;

	*)
		echo "please choose one {dump | restore}"
		exit 1
esac

