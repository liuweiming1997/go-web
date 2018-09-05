/*
* @Author: vimiliu
* @Date:   2018-09-05 14:21:12
* @Last Modified by:   vimiliu
* @Last Modified time: 2018-09-05 15:00:10
 */

package tool

import "github.com/sundayfun/go-web/tool/filter"

func GetTitleFromHtml(html string) string {
	res := filter.ReTitle.FindAllString(html, -1)
	if len(res) == 0 {
		return ""
	}
	ans := res[0][7 : len(res[0])-8]
	return ans
}
