#!/bin/bash

# local test here
function localtest() {
	echo "local debug begin"
	cd ./main-server
	go run ./web-crawler.go
}

case "$1" in
	localtest) 
		localtest
		;;

    *)
        echo $"must choose localtest"
        exit 1
esac