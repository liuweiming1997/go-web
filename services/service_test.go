package services

import (
	"fmt"
	"testing"
)

func TestMarkDown(t *testing.T) {
	res := MarkDownFromTitleAndURL(1, "vimi", "www.baidu.com")
	fmt.Println(res)
}
