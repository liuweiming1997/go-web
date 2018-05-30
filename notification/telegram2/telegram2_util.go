package telegram2

import "fmt"

func formatTelegramMethodUrl(token, method string) string {
	return fmt.Sprintf("%s%s/%s", telegramUrl, token, method)
}

func getDefaultTelegramMsg() *telegramMsg {
	return &telegramMsg{
		Token:     "",
		ChatId:    "",
		Text:      "[default_text](www.baidu.com)",
		ParseMode: "markdown",
		Media:     nil,
		Type:      "text",
		DisableWebPagePreview: false,
		DisableNotification:   false,
		Button:                nil,
	}
}

func (s *telegramMsg) toMessage() *message {
	return &message{
		ChatId:                s.ChatId,
		Text:                  s.Text,
		ParseMode:             s.ParseMode,
		DisableWebPagePreview: s.DisableWebPagePreview,
		DisableNotification:   s.DisableNotification,
	}
}

func (s *telegramMsg) toMessageWithButton() *messageWithButton {
	// make([][]*inlineKeyboardButton, len(s.Button)) will build len(s.Button) nil in element
	replyMarkUp := &inlineKeyboardMarkup{make([][]*inlineKeyboardButton, 0, len(s.Button))}
	button := []*inlineKeyboardButton{}

	for i := 0; i < len(s.Button); i++ {
		button = append(button, &inlineKeyboardButton{s.Button[i][0], s.Button[i][1]})
	}

	replyMarkUp.InlineKeyBoard = append(replyMarkUp.InlineKeyBoard, button)

	return &messageWithButton{
		ChatId:                s.ChatId,
		Text:                  s.Text,
		ParseMode:             s.ParseMode,
		DisableWebPagePreview: s.DisableWebPagePreview,
		DisableNotification:   s.DisableNotification,
		ReplyMarkUp:           replyMarkUp,
	}
}

func (s *telegramMsg) toMediaGroup() *mediaGroup {
	arr := []*mediaArr{}
	for i := 0; i < len(s.Media); i++ {
		arr = append(arr, &mediaArr{
			Type:      s.Type,
			Describe:  "[" + s.Media[i][0] + "]" + "(" + s.Media[i][1] + ")",
			MediaURL:  s.Media[i][1],
			ParseMode: s.ParseMode})
	}
	return &mediaGroup{
		ChatId:              s.ChatId,
		DisableNotification: s.DisableNotification,
		Media:               arr,
	}
}
