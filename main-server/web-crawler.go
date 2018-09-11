package main

import (
	"runtime"

	"github.com/sundayfun/go-web/main-server/servies/cloudyun"
	"github.com/sundayfun/go-web/main-server/servies/niuke"
)

func main() {
	// go cnblogs.Producer()
	go cloudyun.Producer()
	go niuke.Producer()
	for {
		runtime.Gosched()
	}
}
