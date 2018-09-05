package telegram

const (
	telegramUrl    = "https://api.telegram.org/bot"
	sendMessage    = "sendMessage"
	sendMediaGroup = "sendMediaGroup"
	getMe          = "getMe"
	postMethod     = "application/json"
)

//Expose to user
type telegramMsg struct {
	Token                 string      `json:"token"`
	ChatId                string      `json:"chat_id"`
	Text                  string      `json:"text"`
	ParseMode             string      `json:"parse_mode"`
	Media                 [][2]string `json:"media"` // Describe + url
	Type                  string      `json:"type"`  // text, video, photo
	DisableWebPagePreview bool        `json:"disable_web_page_preview"`
	DisableNotification   bool        `json"disable_notification"`
	Button                [][2]string `json:"reply_markup"` // Describe + url
}

// inter use

//**********sendMessage***************

type inlineKeyboardButton struct {
	Text string `json:"text"`
	Url  string `json:"url"`
}

type inlineKeyboardMarkup struct {
	InlineKeyBoard [][]*inlineKeyboardButton `json:"inline_keyboard"`
}

type messageWithButton struct {
	ChatId                string                `json:"chat_id"`
	Text                  string                `json:"text"`
	ParseMode             string                `json:"parse_mode"`
	DisableWebPagePreview bool                  `json:"disable_web_page_preview"`
	DisableNotification   bool                  `json:"disable_notification"`
	ReplyMarkUp           *inlineKeyboardMarkup `json:"reply_markup"`
}

type message struct {
	ChatId                string `json:"chat_id"`
	Text                  string `json:"text"`
	ParseMode             string `json:"parse_mode"`
	DisableWebPagePreview bool   `json:"disable_web_page_preview"`
	DisableNotification   bool   `json:"disable_notification"`
}

//**********sendMessage***************

//************************************
//**********sendMediaGroup************
type mediaArr struct {
	Type      string `json:"type"`
	MediaURL  string `json:"media"`
	Describe  string `json:"caption"`
	ParseMode string `json:"parse_mode"`
}

type mediaGroup struct {
	ChatId              string      `json:"chat_id"`
	DisableNotification bool        `json:"disable_notification"`
	Media               []*mediaArr `json:"media"`
}

//**********sendMediaGroup************
