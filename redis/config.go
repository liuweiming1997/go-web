package redis

import (
	"fmt"
	"os"
	"time"

	"github.com/gomodule/redigo/redis"
	"github.com/sirupsen/logrus"
	"github.com/sundayfun/go-web/tool"
)

var GlobalRedisPool *redis.Pool

func init() {
	addr := os.Getenv(tool.RedisHost)
	if addr == "" {
		addr = "127.0.0.1:6379"
		fmt.Printf("why %s is empty?\n", tool.RedisHost)
	}
	logrus.Infof("%s = %s\n", tool.RedisHost, addr)
	GlobalRedisPool = newPool(addr)
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
