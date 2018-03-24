#!/bin/bash

# local test here
function localtest() {
	echo "local debug begin"
	cd ./server
	go run ./main.go
}

case "$1" in
	localtest) 
		localtest
		;;

    *)
        echo $"must choose localtest"
        exit 1
esac