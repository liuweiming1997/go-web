package main

import (
	"runtime"
  "time"
	"github.com/sundayfun/go-web/main-server/servies/cloudyun"
	"github.com/sundayfun/go-web/main-server/servies/niuke"
	"github.com/sundayfun/go-web/main-server/servies/niukemianjin"
)

func main() {
	go cnblogs.Producer()
	go cloudyun.Producer()
	go niuke.Producer()
	for {
    time.Sleep(time.Second * 30)
		runtime.Gosched()
	}
}
