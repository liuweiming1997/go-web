package telegram

import (
	"fmt"
	"regexp"
	"runtime"
	"sync"
	"time"

	"github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/sirupsen/logrus"
	. "github.com/sundayfun/go-web/services"
	"github.com/sundayfun/go-web/util/filter"
)

var GlobalTelegramBot *telegramBot
var GlobalLock = &sync.Mutex{}

func init() {
	logrus.Infof("the number of cpu: %d", runtime.NumCPU())
	// runtime.GOMAXPROCS(1000)
	GlobalTelegramBot = getTelegramBot()
	GlobalTelegramBot.startChat()
	GlobalTelegramBot.startSpider("https://www.cnblogs.com/", filter.ReHtml, 30, TelegramChatIDGroup)
	GlobalTelegramBot.startNotification()
}

func (s *telegramBot) startNotification() {
	logrus.Info("start notification consumer")
	for {
		chatID := <-s.chatID
		chatMessage := <-s.chatMessage

		msg := tgbotapi.NewMessage(chatID, chatMessage)

		_, err := s.bot.Send(msg)
		if err != nil {
			logrus.Error(err)
		}
	}
}

func (s *telegramBot) startSpider(url string, re *regexp.Regexp, updateTimePerMinute int64, chatID int64) {
	//TODO: same url check use redis
	logrus.Info("start spider consumer")
	ticker := time.NewTicker(time.Duration(updateTimePerMinute) * time.Minute)
	go func() {
		for t := range ticker.C {
			fmt.Println(t.String())
			html, err := HtmlFromUrl(url)
			if err != nil {
				logrus.Error(err)
				continue
			}
			chatMessage := t.String()[0:19] + "\n"
			chatMessage += WannerFromRegexp(re, html)
			if len(chatMessage) == 20 {
				continue
			}
			PushMessageToTelegram(chatID, chatMessage)
		}
	}()
}

func (s *telegramBot) startChat() {
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
