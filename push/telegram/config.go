package telegram

import (
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/kelseyhightower/envconfig"
	"github.com/sirupsen/logrus"
)

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

type tokenStruct struct {
	Token string `default:""`
}

func getTelegramBot() *telegramBot {
	var s tokenStruct
	err := envconfig.Process("Telegram_Token", &s)
	if err != nil {
		logrus.Fatal(err)
	}

	_token := s.Token

	logrus.Info("Telegram_Token:", _token)

	_bot, err := tgbotapi.NewBotAPI(_token)
	if err != nil {
		logrus.Fatal(err)
	}
	logrus.Infof("Authorized on account %s", _bot.Self.UserName)

	return &telegramBot{
		bot:   _bot,
		token: _token,
		chat:  make(chan *Message, 100),
		// chatID:      make(chan int64, 100),
		// chatMessage: make(chan string, 100),
	}
}
