package main

import (
	"runtime"
	"time"
	"github.com/sundayfun/go-web/main-server/servies/cloudyun"
	"github.com/sundayfun/go-web/main-server/servies/cnblogs"
	"github.com/sundayfun/go-web/main-server/servies/dzone"
	"github.com/sundayfun/go-web/main-server/servies/niuke"
	"github.com/sundayfun/go-web/main-server/servies/niukemianjin"
)

func main() {
	go cnblogs.Producer()
	go cloudyun.Producer()
	go niuke.Producer()
  go niukemianjin.Producer()
  go dzone.Producer()
	for {
		time.Sleep(time.Second * 30)
		runtime.Gosched()
	}
}
