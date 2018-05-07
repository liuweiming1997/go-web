package push

import (
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/sirupsen/logrus"
)

const (
	VimiBotKey          = ""
	TelegramChatIDGroup = -264517585
	TelegramChatIDVimi  = 505481672
)

var (
	vimiBot *tgbotapi.BotAPI
)

func init() {
	//can not use vimiBot, err := tgbotapi.NewBotAPI(TelegramKey)
	//cause global_vimiBot will be nil
	var err error
	vimiBot, err = tgbotapi.NewBotAPI(VimiBotKey)
	if err != nil {
		logrus.Fatal(err)
	}
	logrus.Infof("Authorized on account %s", vimiBot.Self.UserName)
}

func PushMessageToTelegram(b []byte) error {
	msg := tgbotapi.NewMessage(TelegramChatIDGroup, string(b))
	res, err := vimiBot.Send(msg)
	if err != nil {
		return err
	}
	logrus.Debug(res)
	return nil
}

func StateChat() {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := vimiBot.GetUpdatesChan(u)
	if err != nil {
		logrus.Fatal(err)
	}
	for update := range updates {
		if update.Message == nil {
			continue
		}

		logrus.Infof("[%s] %s", update.Message.From.UserName, update.Message.Text)

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
		msg.ReplyToMessageID = update.Message.MessageID

		vimiBot.Send(msg)
	}
}
