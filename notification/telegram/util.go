package telegram

func PushMessageToTelegram(chatID int64, chatMessage string) {
	msg := &Message{
		chatID:      chatID,
		chatMessage: chatMessage,
	}
	GlobalTelegramBot.chat <- msg
	// GlobalTelegramBot.chatID <- chatID
	// GlobalTelegramBot.chatMessage <- chatMessage
}
