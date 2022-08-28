/*
* @Author: vimiliu
* @Date:   2018-09-05 13:45:48
* @Last Modified by:   vimiliu
* @Last Modified time: 2018-09-06 14:03:43
 */

/**
 * 爬取博客园
 */
package cnblogs

import (
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
	indexUrl   = "https://www.cnblogs.com/"
	updateTime = 3 // s
	info       = "爬虫--博客园"
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

			title := tool.GetTitleFromHtml(urlHtml)

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

	// 再按照固定的东西筛选
	t := &filter.VimiRegexp{
		BeginWith:   []string{},
		MustContain: []string{`go`, `python3`, `jenkins`, `单元测试`, `unittest`, `测试`, `bazel`, `nat`, `tcp`, `网络`, `udp`, `协程`, `路由`},
		EndWith:     []string{},
	}
	re := t.GetRegexp()
	finalResult := []string{}
	for _, url := range urls {
		urlHtml, err := tool.GetHtmlFromUrl(url)
		if err != nil {
			logrus.Info(info + " --- " + err.Error())
			continue
		}

		if re.MatchString(tool.GetTitleFromHtml(urlHtml)) {
			finalResult = append(finalResult, url)
		}
	}
	return finalResult
}

func getAllUrlFromHtml(html string) []string {
	t := &filter.VimiRegexp{
		BeginWith:   []string{`http://`, `https://`},
		MustContain: []string{},
		EndWith:     []string{`.html`},
	}
	re := t.GetRegexp()
	finalResult := []string{}
	urls := re.FindAllString(html, -1)
	for _, url := range urls {
		hashValue := strconv.FormatUint(tool.String2uint64(url), 10)
		if !redis.Exist([]byte(hashValue)) {
			redis.Set([]byte(hashValue), []byte(tool.UrlKey))
			finalResult = append(finalResult, url)
		}
	}
	return finalResult
}
