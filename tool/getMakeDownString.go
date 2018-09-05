/*
* @Author: vimiliu
* @Date:   2018-09-05 14:24:24
* @Last Modified by:   vimiliu
* @Last Modified time: 2018-09-05 19:55:11
 */
package tool

func GetMarkDownString(title string, url string) string {
	// fmt.Printf("title = %s\nurl = %s\n\n", title, url)
	ans := "[" + title + "]"
	ans += "(" + url + ")"
	return ans
}
