package telegram

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"

	"github.com/sirupsen/logrus"
)

func telegramAuth(token string) bool {
	url := formatTelegramMethodUrl(token, getMe)
	resp, err := http.Get(url)
	if err != nil {
		logrus.Error(err)
		return false
	}
	defer resp.Body.Close()
	return resp.StatusCode == 200
}

func TelegramPush(msg *telegramMsg) error {
	if msg.Token == "" {
		return errors.New("empty token")
	}
	if msg.ChatId == "" {
		return errors.New("empty chat_id")
	}
	method := "sendMessage"
	jsonValue := []byte{}
	var err error

	switch msg.Type {
	case "text":
		if msg.Button != nil {
			jsonValue, err = json.Marshal(msg.toMessageWithButton())
		} else {
			jsonValue, err = json.Marshal(msg.toMessage())
		}
	default:
		method = "sendMediaGroup"
		jsonValue, err = json.Marshal(msg.toMediaGroup())
	}

	if err != nil {
		return err
	}
	url := formatTelegramMethodUrl(msg.Token, method)
	jsonValueWithIOReader := bytes.NewBuffer(jsonValue)
	resp, err := http.Post(url, postMethod, jsonValueWithIOReader)

	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		b, _ := ioutil.ReadAll(resp.Body)
		return errors.New(string(b))
	}
	return nil
}
