package main

import (
	"strconv"
	"strings"

	. "github.com/sundayfun/go-web/push/telegram"
	"github.com/sundayfun/go-web/services"
	. "github.com/sundayfun/go-web/util/filter"
)

var mp = make(map[string]bool)

func main() {
	searchJianShu("https://www.jianshu.com/c/0f5cb8eb7927", 60)
	searchJianShu("https://www.jianshu.com/c/3e489dead7a7", 75)
	searchJianShu("https://www.jianshu.com/c/20f7f4031550?utm_medium=index-collections&utm_source=desktop", 30)

	searchBoKeYuan("https://www.cnblogs.com/", 30, TelegramChatIDGroup)

	GlobalTelegramBot.StartChat()
	GlobalTelegramBot.StartNotification()
	for {

	}
	return
}

func searchJianShu(URL string, dis int64) {
	s := &VimiRegexp{
		BeginWith:   []string{`href="`},
		EndWith:     []string{`"`},
		MustContain: []string{`/p`},
	}

	re := s.GetRegexp()

	WannerFromHtml := func(text string) string {
		jianshu := "https://www.jianshu.com"
		str := re.FindAllString(text, -1)
		ans := ""
		id := 0
		for _, val := range str {
			url := jianshu + val[6:len(val)-1]
			if mp[url] {
				continue
			}
			mp[url] = true

			if strings.Contains(url, "#comments") {
				continue
			}
			//TODO: change regexp to mustcontain
			title := services.TitleFromUrl(url, ReTitle)
			if title == "" {
				continue
			}

			if strings.Contains(title, "404 Not Found") {
				continue
			}

			ans += title + "\n"
			ans += strconv.Itoa(id) + " "
			ans += url + "\n"
			ans += "\n"
			id++
		}
		return ans
	}
	GlobalTelegramBot.StartSpider(URL, WannerFromHtml, dis, TelegramChatIDVimi)
}

func searchBoKeYuan(URL string, dis int64, chatID int64) {
	s := &VimiRegexp{
		BeginWith:   []string{"http://", "https://"},
		EndWith:     []string{".html"},
		MustContain: []string{""},
	}
	re := s.GetRegexp()

	WannerFromHtml := func(html string) string {
		str := re.FindAllString(html, -1)
		id := 0
		ans := ""
		for _, val := range str {
			if mp[val] {
				continue
			}
			title := services.TitleFromUrl(val, ReTitle)
			if title == "" {
				continue
			}
			mp[val] = true
			ans += title + "\n"
			ans += strconv.Itoa(id) + " "
			ans += val + "\n"
			ans += "\n"
			id++
		}
		return ans
	}
	GlobalTelegramBot.StartSpider(URL, WannerFromHtml, dis, chatID)
}
