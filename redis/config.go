package redis

import (
	"fmt"
	"os"
	"time"

	"github.com/gomodule/redigo/redis"
	"github.com/sundayfun/go-web/tool"
)

var GlobalRedisPool *redis.Pool

func init() {
	addr := os.Getenv(tool.RedisHost)
	if addr == "" {
		fmt.Printf("why %s is empty?", tool.RedisHost)
	}
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
				for {
					fmt.Println("do this")
				}
				_, err := c.Do("PING")
				return err
			}
			return nil
		},
	}
}
