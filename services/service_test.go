package services

import (
	"fmt"
	"testing"

	"github.com/sundayfun/go-web/tool/filter"
)

func TestMarkDown(t *testing.T) {
	res := TitleFromUrl(`https://www.cnblogs.com/xyblogs/p/9090199.html`, filter.ReTitle)
	fmt.Println(res)
}
