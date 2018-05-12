package services

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strconv"

	"github.com/sundayfun/go-web/util/filter"
)

func TitleFromUrl(url string) string {
	html, _ := HtmlFromUrl(url)
	res := filter.ReTitle.FindAllString(html, -1)
	if len(res) == 0 {
		return ""
	}
	fmt.Println(res)
	ans := res[0][7 : len(res[0])-8]
	return ans
}

func HtmlFromUrl(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

var mp = make(map[string]bool)

func WannerFromRegexp(re *regexp.Regexp, text string) string {
	message := ""
	str := re.FindAllString(text, -1)
	id := 1
	for _, val := range str {
		if mp[val] {
			continue
		}
		mp[val] = true
		title := TitleFromUrl(val)
		message += title + "\n"
		message += strconv.Itoa(id) + " "
		message += val + "\n"
		message += "\n"
		id++
	}
	return message
}
