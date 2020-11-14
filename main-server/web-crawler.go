package main

import (
	"runtime"
	"time"
	"github.com/sundayfun/go-web/main-server/servies/cloudyun"
	"github.com/sundayfun/go-web/main-server/servies/cnblogs"
	"github.com/sundayfun/go-web/main-server/servies/drdr"
	"github.com/sundayfun/go-web/main-server/servies/niuke"
	"github.com/sundayfun/go-web/main-server/servies/niukemianjin"
	"github.com/sundayfun/go-web/main-server/servies/draveness"
)

func main() {
	go cnblogs.Producer()
	go cloudyun.Producer()
	go niuke.Producer()
  go niukemianjin.Producer()
  go drdr.Producer()
  go draveness.Producer()

	for {
		time.Sleep(time.Second * 30)
		runtime.Gosched()
	}
}
