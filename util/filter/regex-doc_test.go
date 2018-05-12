package filter

import (
	"fmt"
	"testing"
)

const (
	key = ` https://static.zhihu.com/heifetz/main.question-routes.111fc68095704bf5f2c0.js"></script></body></html`
)

func TestReImageURL(t *testing.T) {
	str := ReImageURL.FindAllString(key, -1)
	for i, val := range str {
		fmt.Println(i, val)
	}
}

func TestReTitle(t *testing.T) {
	key := `<sdf><title>谈谈我开发过的几套语音通信解决方案 - davidtym - 博客园</title>sadfsadf`
	str := ReTitle.FindAllString(key, -1)
	for i, val := range str {
		fmt.Println(i, val)
	}
}
