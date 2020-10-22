#!/bin/bash
export GOPROXY=https://goproxy.io
export GO111MODULE=on

# please make sure running this script using ./script/setup.sh
go build -o main main-server/web-crawler.go
# go run main-server/web-crawler.go
