package telegram

import "testing"

func TestPushMarkDown(t *testing.T) {
	PushMessageToTelegram(TelegramChatIDGroup, "test")
	GlobalTelegramBot.StartNotification()
}
