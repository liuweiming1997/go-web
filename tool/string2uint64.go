/*
* @Author: vimiliu
* @Date:   2018-09-05 14:08:32
* @Last Modified by:   vimiliu
* @Last Modified time: 2018-09-05 14:43:14
 */

// string hash
package tool

const (
	seed = 131
)

func String2uint64(str string) uint64 {
	var ans uint64 = 0
	for i := 0; i < len(str); i++ {
		ans = ans*seed + uint64(str[i])
	}
	return ans
}
