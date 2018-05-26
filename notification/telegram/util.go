package telegram

func PushMessageToTelegram(chatID int64, chatMessage string) {
	if chatMessage == "" {
		return
	}
	msg := &Message{
		chatID:      chatID,
		chatMessage: chatMessage,
	}
	GlobalTelegramBot.chat <- msg
	// GlobalTelegramBot.chatID <- chatID
	// GlobalTelegramBot.chatMessage <- chatMessage
}
