package main

import (
	"runtime"

	"github.com/sundayfun/go-web/main-server/servies/cloudyun"
	"github.com/sundayfun/go-web/main-server/servies/cnblogs"
)

func main() {
	go cnblogs.Producer()
	go cloudyun.Producer()
	for {
		runtime.Gosched()
	}
}
