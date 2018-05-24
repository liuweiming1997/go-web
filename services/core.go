package services

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strconv"

	"github.com/sirupsen/logrus"
	"github.com/sundayfun/go-web/redis"
	"github.com/sundayfun/go-web/tool"
	"github.com/sundayfun/go-web/tool/filter"
)

func TitleFromUrl(url string, re *regexp.Regexp) string {
	html, err := HtmlFromUrl(url)
	if err != nil {
		return ""
	}
	res := re.FindAllString(html, -1)
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

	if resp.StatusCode != 200 {
		return "", fmt.Errorf("statusCode wanner 200 but have %d", resp.StatusCode)
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func WannerFromRegexp(re *regexp.Regexp, text string) string {
	message := ""
	str := re.FindAllString(text, -1)
	id := 1
	for _, val := range str {
		if redis.Exist([]byte(val)) {
			continue
		}
		redis.Set([]byte(val), []byte(tool.UrlKey))
		logrus.Debug(val)
		title := TitleFromUrl(val, filter.ReTitle)
		if title == "" {
			continue
		}
		message += title + "\n"
		message += strconv.Itoa(id) + " "
		message += val + "\n"
		message += "\n"
		id++
	}
	return message
}

func TestFunc(text string) string {
	return text
}
