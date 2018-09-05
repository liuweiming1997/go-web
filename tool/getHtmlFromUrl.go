package tool

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

/*
* @Author: vimiliu
* @Date:   2018-09-05 14:20:20
* @Last Modified by:   vimiliu
* @Last Modified time: 2018-09-05 14:20:48
 */

func GetHtmlFromUrl(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return "", fmt.Errorf("statusCode wanner 200 but have %d", resp.StatusCode)
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(data), nil
}
