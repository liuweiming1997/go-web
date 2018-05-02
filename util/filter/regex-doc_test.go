package filter

import (
	"fmt"
	"testing"
)

const (
	key = `<img src="http://img01.tooopen.com/Downs/thumb/sy_201101291355250450-1.jpg" onerror="javascript:this.src='http://resource.tooopen.com/image/no-img-192.gif';this.onerror=null;" alt="韩国模特“天海”" class="imgItem"/>`
)

func TestReImageURL(t *testing.T) {
	str := ReImageURL.FindAllString(key, -1)
	for i, val := range str {
		fmt.Println(i, val)
	}
}
