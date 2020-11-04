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
	client := &http.Client{}
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", err
	}
	request.Header.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.75 Safari/537.36")

	response, err := client.Do(request)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()

	if response.StatusCode != 200 {
		return "", fmt.Errorf("statusCode wanner 200 but have %d", response.StatusCode)
	}

	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", err
	}
	return string(data), nil
}
