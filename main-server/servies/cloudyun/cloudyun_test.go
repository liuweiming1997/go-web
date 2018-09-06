/*
* @Author: vimiliu
* @Date:   2018-09-06 10:17:53
* @Last Modified by:   vimiliu
* @Last Modified time: 2018-09-06 13:11:35
 */

package cloudyun

import (
	"testing"
)

const (
	testString = "httpfsdfsdhttp://fuckme.htmlfsjdlkf"
)

// func TestB(f *testing.T) {
// 	t := &filter.VimiRegexp{
// 		BeginWith:   []string{`http://`, `https://`},
// 		MustContain: []string{},
// 		EndWith:     []string{`.html`},
// 	}
// 	re := t.GetRegexp()
// 	urls := re.FindAllString(testString, -1)
// 	for _, url := range urls {
// 		fmt.Println(url)
// 	}
// }

func TestA(t *testing.T) {
	Producer()
}
