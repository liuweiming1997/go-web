/*
* @Author: vimiliu
* @Date:   2018-09-11 22:00:33
* @Last Modified by:   vimiliu
* @Last Modified time: 2018-09-11 23:10:22
 */

package niuke

import (
	"fmt"
	"strconv"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/sundayfun/go-web/env"
	"github.com/sundayfun/go-web/notification/telegram"
	"github.com/sundayfun/go-web/redis"
	"github.com/sundayfun/go-web/tool"
	"github.com/sundayfun/go-web/tool/filter"
)

const (
	indexUrl   = "https://www.nowcoder.com/activity/campus2019"
	baseUrl    = "https://www.nowcoder.com"
	updateTime = 300 // s
	info       = "爬虫--牛客网"
)

func Producer() {
	logrus.Info(info)
	ticker := time.NewTicker(time.Duration(updateTime) * time.Second)
	for t := range ticker.C {
		logrus.Info(info + " --- " + t.String())
		html, err := tool.GetHtmlFromUrl(indexUrl)
		if err != nil {
			logrus.Info(info + " --- " + err.Error())
			continue
		}
		urls := getWantFromHtml(html)
		for _, url := range urls {
			urlHtml, _ := tool.GetHtmlFromUrl(url)

			title := getTitleFromHtml(urlHtml)
			fmt.Println(title, url)
			msg := telegram.GetDefaultTelegramMsg()
			msg.Token = env.GetTelegramToken()
			msg.ChatId = telegram.TelegramChatIDGroup
			if msg.Token == "" {
				msg.Token = "516690928:AAH4l2EyC8YAFalLut6ZMoWv-1BrqgoAkfo"
			}
			msg.Text = tool.GetMarkDownString(title, url)
			err = telegram.TelegramPush(msg)
			if err != nil {
				logrus.Info(info + " --- " + err.Error())
			}
		}
	}
}

func getWantFromHtml(html string) []string {
	// 先找出所有url
	urls := getAllUrlFromHtml(html)
	return urls
}

func getAllUrlFromHtml(html string) []string {
	t := &filter.VimiRegexp{
		BeginWith:   []string{`/activity/campus2019`},
		MustContain: []string{},
		EndWith:     []string{`"`},
	}
	re := t.GetRegexp()
	finalResult := []string{}
	urls := re.FindAllString(html, -1)
	for _, url := range urls {
		url = baseUrl + url[:len(url)-1]
		hashValue := strconv.FormatUint(tool.String2uint64(url), 10)
		if !redis.Exist([]byte(hashValue)) {
			redis.Set([]byte(hashValue), []byte(tool.UrlKey))
			finalResult = append(finalResult, url)
		}
	}
	return finalResult
}

func getTitleFromHtml(html string) string {
	re := (&filter.VimiRegexp{
		BeginWith:   []string{`com-title`},
		MustContain: []string{},
		EndWith:     []string{`</h2>`},
	}).GetRegexp()

	urls := re.FindAllString(html, -1)
	if len(urls) == 0 {
		return "errorPage"
	}
	// <h2 class="com-title">海康威视2019校招</h2>
	return urls[0][11 : len(urls[0])-5]
}
