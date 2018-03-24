#!/bin/bash

DBUSER=root
DBHOST=127.0.0.1
DBNAME=homework
DBPASSWORD=dc

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

