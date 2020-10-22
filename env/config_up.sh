#!/bin/bash

set -e

export MYSQL_ROOT_PASSWORD=vimi
export MYSQL_DATABASE=homework
export VIMI_DB_HOST=127.0.0.1
export Telegram_Token="your Telegram_Token"
export Redis_Host=127.0.0.1:6379
# export HTTP_PROXY=socks5://sslocal:1080
# export HTTPS_PROXY=socks5://sslocal:1080

export GOPROXY=https://goproxy.io
export GO111MODULE=on
