#!/bin/bash

set -e

# local test here
function localtest() {
	echo "local debug begin"
	./env/config_up.sh
	cd ./main-server
	go run ./web-crawler.go
	cd ..
	./env/config_down.sh
}

case "$1" in
	localtest) 
		localtest
		;;

    *)
        echo $"must choose localtest"
        exit 1
esac