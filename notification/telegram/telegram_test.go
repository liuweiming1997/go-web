package telegram

import (
	"fmt"
	"testing"
)

func TestAuth(t *testing.T) {
	token := "516690928:AAH4l2EyC8YAFalLut6ZMoWv-1BrqgoAkfo"
	fmt.Println(telegramAuth(token))
}

func TestTelegramPush(t *testing.T) {
	res := GetDefaultTelegramMsg()
	res.ChatId = "505481672"
	res.Token = "516690928:AAH4l2EyC8YAFalLut6ZMoWv-1BrqgoAkfo"
	button := [][2]string{[2]string{"vimi", "www.baidu.com"}, [2]string{"hello", "http://www.facebook.com"}}
	res.Button = button

	// jsonValue, _ := json.Marshal(res.toMessageWithButton())
	// fmt.Printf("%+v", string(jsonValue))

	err := TelegramPush(res)
	if err != nil {
		fmt.Println(err)
	}
}

func TestTelegramPushMediaGroup(t *testing.T) {
	res := GetDefaultTelegramMsg()
	res.ChatId = "505481672"
	res.Token = "516690928:AAH4l2EyC8YAFalLut6ZMoWv-1BrqgoAkfo"
	res.Type = "video"

	res.Media = [][2]string{[2]string{"tian", "https://www.google.com/imgres?imgurl=http%3A%2F%2Fhbimg.b0.upaiyun.com%2Ff18520e514ec5e8345216af5940a371d3d07efae233e5-MCeudN_fw658&imgrefurl=http%3A%2F%2Fhuaban.com%2Fpins%2F1341947953%2F&docid=nD_px8fenqYtgM&tbnid=nGpxFu6IY5xhaM%3A&vet=10ahUKEwj6p_Hy4ajbAhWqGDQIHUBLDYUQMwg7KAgwCA..i&w=443&h=562&bih=649&biw=1301&q=%E5%A4%A9%E6%B5%B7%E7%BF%BC&ved=0ahUKEwj6p_Hy4ajbAhWqGDQIHUBLDYUQMwg7KAgwCA&iact=mrc&uact=8"}}

	err := TelegramPush(res)
	if err != nil {
		fmt.Println(err)
	}
}
