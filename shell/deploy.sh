#!/bin/bash

#39.108.100.86
server_address=119.23.231.141

function deploy() {
	#  rsync的desc会自动创建一个目录，所以这样就是/root/go-web
	echo "maybe a little bit slow because will push this file to your-server"
	rsync -avz --delete ../go-web root@${server_address}:/root
	# ssh root@${server_address} "cd go-web/docker; 
 #                                   docker-compose up --build -d db; 
 #                                   docker-compose up --build -d dbrestore; 
 #                                   docker-compose up --build -d server;"
 	ssh root@${server_address} "cd go-web/docker;
                                   docker-compose up --build -d server;"
}

function dbrestore() {
	mysql -h $VIMI_DB_HOST -uroot -p$MYSQL_ROOT_PASSWORD $MYSQL_DATABASE < ./db/sql/latest_dump.sql
}

function now() {
	echo pwd;
}

case "$1" in
	deploy) 
		deploy
		;;

	dbrestore)
        dbrestore
        ;;
     now)
		now
		;;
    *)
        echo $"must choose deploy | dbrestore"
        exit 1
esac