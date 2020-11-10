package drdr

import (
	"strconv"
	"time"

	"github.com/axgle/mahonia"
	"github.com/sirupsen/logrus"
	"github.com/sundayfun/go-web/env"
	"github.com/sundayfun/go-web/notification/telegram"
	"github.com/sundayfun/go-web/redis"
	"github.com/sundayfun/go-web/tool"
	"github.com/sundayfun/go-web/tool/filter"
)

const (
	indexUrl   = "https://drmingdrmer.github.io/"
	updateTime = 300 // s
	info       = "爬虫--drdr"
)

func Producer() {
	logrus.Info(info + " --- " + "start")
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
			title := ConvertToString(tool.GetTitleFromHtml(urlHtml), "utf-8", "utf-8")

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

func ConvertToString(src string, srcCode string, tagCode string) string {
	srcCoder := mahonia.NewDecoder(srcCode)
	srcResult := srcCoder.ConvertString(src)
	tagCoder := mahonia.NewDecoder(tagCode)
	_, cdata, _ := tagCoder.Translate([]byte(srcResult), true)
	result := string(cdata)
	return result
}

func getWantFromHtml(html string) []string {
	// 先找出所有url
	urls := getAllUrlFromHtml(html)
	finalResult := make([]string, len(urls))
	for idx:= 0; idx < len(urls); idx++ {
		finalResult[idx] = indexUrl + urls[idx][6:]
	}
	return finalResult
	// return finalResult
}

func getAllUrlFromHtml(html string) []string {
	t := &filter.VimiRegexp{
		BeginWith:   []string{`href`},
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
