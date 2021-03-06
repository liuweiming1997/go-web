package services

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"

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
		message += MarkDownFromTitleAndURL(title, val)
		message += "\n\n"
	}
	return message
}

func MarkDownFromTitleAndURL(title string, url string) string {
	fmt.Printf("title = %s\nurl = %s\n\n", title, url)
	ans := "[" + title + "]"
	ans += "(" + url + ")"
	return ans
}
