package main

import (
	"runtime"
	"time"

	"github.com/sundayfun/go-web/main-server/servies/cnblogs"
	"github.com/sundayfun/go-web/main-server/servies/niuke"
)

func main() {
	go cnblogs.Producer()
	go niuke.Producer()
	for {
		time.Sleep(time.Second * 30)
		runtime.Gosched()
	}
}
