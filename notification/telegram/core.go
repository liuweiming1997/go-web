package telegram

import (
	"fmt"
	"time"

	"github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/sirupsen/logrus"
	"github.com/sundayfun/go-web/redis"
	. "github.com/sundayfun/go-web/services"
	"github.com/sundayfun/go-web/tool"
	"github.com/sundayfun/go-web/tool/filter"
)

var done = make(chan bool, 1)

func (s *telegramBot) StartNotification() {
	logrus.Info("start notification consumer")
	for {
		temp := <-GlobalTelegramBot.chat
		msg := tgbotapi.NewMessage(temp.chatID, temp.chatMessage)
		msg.ParseMode = "markdown"
		msg.DisableWebPagePreview = true
		_, err := s.bot.Send(msg)
		if err != nil {
			logrus.Error(err)
		}
	}
}

func (s *telegramBot) StartSpider(url string, WannerFromHtml func(string) string, updateTimePerMinute int64, chatID int64) {

	logrus.Info("start spider consumer")
	ticker := time.NewTicker(time.Duration(updateTimePerMinute) * time.Second)
	go func() {
		for t := range ticker.C {
			fmt.Println("now time = " + t.String())
			html, err := HtmlFromUrl(url)
			if err != nil {
				logrus.Error("doing spider ", err)
				continue
			}
			chatMessage := WannerFromHtml(html)
			PushMessageToTelegram(chatID, chatMessage)
		}
	}()
	// ok := make(chan bool, 1)
	// ok <- true
	// go func() {
	// 	for t := range ok {
	// 		fmt.Println("what is it? ", t)
	// 		html, err := HtmlFromUrl(url)
	// 		if err != nil {
	// 			logrus.Error("doing spider ", err)
	// 			continue
	// 		}
	// 		chatMessage := WannerFromHtml(html)
	// 		PushMessageToTelegram(chatID, chatMessage)
	// 		select {
	// 		case <-done:
	// 			break
	// 		case <-time.After(time.Minute):
	// 			logrus.Error("done time out")
	// 			break
	// 		}

	// 		time.Sleep(time.Duration(updateTimePerMinute) * time.Minute)
	// 		ok <- true
	// 	}
	// }()
}

func (s *telegramBot) StartChat() {
	logrus.Info("start chat consumer")
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates, err := s.bot.GetUpdatesChan(u)

	if err != nil {
		logrus.Fatal(err)
	}

	go func() {
		for val := range updates {
			if val.Message == nil {
				continue
			}

			//TOTEST: I save it in redis and check it
			redis.Set([]byte(val.Message.Text), []byte(tool.UrlKey))

			html, err := HtmlFromUrl(val.Message.Text)
			if err != nil {
				PushMessageToTelegram(val.Message.Chat.ID, "it seem that is not a website")
				continue
			}
			text := WannerFromRegexp(filter.ReHtml, html)
			if text == "" {
				PushMessageToTelegram(val.Message.Chat.ID, "no update")
				continue
			}
			PushMessageToTelegram(val.Message.Chat.ID, text)
		}
	}()
}
