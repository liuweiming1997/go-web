package main

import (
	"runtime"

	"github.com/sundayfun/go-web/main-server/servies/cnblogs"
)

func main() {
	go cnblogs.Producer()
	for {
		runtime.Gosched()
	}
}
