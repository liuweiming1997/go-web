package redis

import (
	"fmt"
	"testing"
)

func TestMain(t *testing.T) {
	conn1 := GlobalRedisPool.Get()
	conn2 := GlobalRedisPool.Get()
	fmt.Println(conn1, conn2, conn3)
}
