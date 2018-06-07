package telegram

import (
	"fmt"
	"os"
	"runtime"

	"github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/sirupsen/logrus"
	"github.com/sundayfun/go-web/tool"
)

var GlobalTelegramBot *telegramBot

func init() {
	logrus.Info(os.Getenv("http_proxy"))
	logrus.Infof("the number of cpu: %d", runtime.NumCPU())
	// runtime.GOMAXPROCS(1000)
	GlobalTelegramBot = getTelegramBot()
}

const (
	TelegramChatIDGroup = -264517585
	TelegramChatIDVimi  = 505481672
)

type Message struct {
	chatID      int64
	chatMessage string
}

type telegramBot struct {
	bot   *tgbotapi.BotAPI `default: nil`
	token string           `default: nil`
	chat  chan *Message    `default: nil`
	// chatID      chan int64  `default: nil`
	// chatMessage chan string `default: nil`
}

func getTelegramBot() *telegramBot {
	token := os.Getenv(tool.TelegramToken)
	if token == "" {
		token = "516690928:AAH4l2EyC8YAFalLut6ZMoWv-1BrqgoAkfo"
		fmt.Printf("why %s is empty?\n", tool.TelegramToken)
	}

	logrus.Info("Telegram_Token:", token)

	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		logrus.Fatal(err)
	}
	logrus.Infof("Authorized on account %s", bot.Self.UserName)

	return &telegramBot{
		bot:   bot,
		token: token,
		chat:  make(chan *Message, 100),
		// chatID:      make(chan int64, 100),
		// chatMessage: make(chan string, 100),
	}
}
