package push

import (
	"github.com/logrus"
	"gopkg.in/telegram-bot-api.v4"
)

const (
	VimiBotKey          = "516690928:AAGrWIVWj3LuUMicwFaajnyt4Z0j-CP6X7U"
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
