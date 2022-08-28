/*
* @Author: vimiliu
* @Date:   2018-09-06 13:13:49
* @Last Modified by:   vimiliu
* @Last Modified time: 2018-09-06 13:47:51
 */

package tool

import (
	"fmt"
	"testing"

	"github.com/axgle/mahonia"
	"github.com/sundayfun/go-web/tool"
)

// 查询云风的blog，然后把gbk 转换成 utf-8

func TestGetHtmlFromUrl(t *testing.T) {
	url := "https://blog.codingnow.com/2018/08/lockstep.html#more"
	html, err := tool.GetHtmlFromUrl(url)
	if err != nil {
		fmt.Println(err)
	}
	html = ConvertToString(html, "gbk", "utf-8")
	fmt.Println("do")
	fmt.Println(html)
}

func ConvertToString(src string, srcCode string, tagCode string) string {
	srcCoder := mahonia.NewDecoder(srcCode)
	srcResult := srcCoder.ConvertString(src)
	tagCoder := mahonia.NewDecoder(tagCode)
	_, cdata, _ := tagCoder.Translate([]byte(srcResult), true)
	result := string(cdata)
	return result
}
