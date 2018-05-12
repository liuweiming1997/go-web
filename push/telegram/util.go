package telegram

func PushMessageToTelegram(chatID int64, chatMessage string) {
	GlobalTelegramBot.chatID <- chatID
	GlobalTelegramBot.chatMessage <- chatMessage
}
