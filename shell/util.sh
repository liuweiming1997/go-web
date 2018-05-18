#!/bin/bash

server_address=95.163.202.160
function getRemote() {
	echo "getting....."
	rsync -avz --delete root@${server_address}:/root/go-web ../
}

function stopRemote() {
	echo "stop....."
	ssh root@${server_address} "docker stop docker_server_1"
}

function logRemote() {
	echo "docker logs -f docker_server_1......"
	ssh root@${server_address} "docker logs -f docker_server_1"
}

case "$1" in
	getRemote) 
		getRemote
		;;

	stopRemote)
		stopRemote
		;;

	logRemote)
		logRemote
		;;
	*)
        echo $"must choose getRemote | stopRemote"
        exit 1
esac