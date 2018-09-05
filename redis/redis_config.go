package redis

import (
	"fmt"
	"time"

	"github.com/gomodule/redigo/redis"
	"github.com/sirupsen/logrus"
	"github.com/sundayfun/go-web/env"
	"github.com/sundayfun/go-web/tool"
)

var GlobalRedisPool *redis.Pool

func init() {
	addr := env.GetRedisHost()
	if addr == "" {
		addr = "127.0.0.1:6379"
		fmt.Printf("why %s is empty?\n", tool.RedisHost)
	}
	logrus.Infof("%s = %s", tool.RedisHost, addr)
	GlobalRedisPool = newPool(addr)

	// go func() {
	// 	ticker := time.NewTicker(time.Hour * 36)
	// 	for t := range ticker.C {
	// 		logrus.Info("time := ", t, "clear all")
	// 		ClearAll()
	// 	}
	// }()
}

func newPool(addr string) *redis.Pool {
	return &redis.Pool{
		MaxIdle: 3,

		IdleTimeout: 240 * time.Second,

		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", addr)
		},

		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			//TODO: find how to use it
			if time.Since(t) > 20*time.Second {
				fmt.Println("do this")
				_, err := c.Do("PING")
				return err
			}
			return nil
		},
	}
}
