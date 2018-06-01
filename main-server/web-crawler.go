package main

import (
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/sirupsen/logrus"
	. "github.com/sundayfun/go-web/notification/telegram"
	"github.com/sundayfun/go-web/redis"
	"github.com/sundayfun/go-web/services"
	"github.com/sundayfun/go-web/tool"
	. "github.com/sundayfun/go-web/tool/filter"
)

func main() {

	conn := redis.GlobalRedisPool.Get()
	defer conn.Close()

	_, err := conn.Do("FLUSHALL")

	if err != nil {
		logrus.Error(err)
	}
	go func() {
		sigs := make(chan os.Signal, 1)
		signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
		sig := <-sigs
		fmt.Println("error -----", sig, "---------")
		PushMessageToTelegram(TelegramChatIDVimi, "something stop me ......")
		os.Exit(tool.ExitCode)
	}()

	//TOTEST checkout redis
	key := "https://www.cnblogs.com/"
	res := redis.Exist([]byte(key))
	if res {
		fmt.Printf("redis_test key = %s : exist\n", key)
	} else {
		fmt.Printf("redis_test key = %s : not exist\n", key)
	}

	searchBoKeYuan("https://www.cnblogs.com", 300, TelegramChatIDGroup)
	searchBole("http://blog.jobbole.com/category/humor-comic/", 300, TelegramChatIDGroup)
	GlobalTelegramBot.StartChat()
	GlobalTelegramBot.StartNotification()
	for {

	}
	return
}

func searchBole(URL string, dis int64, chatID int64) {
	s := &VimiRegexp{
		BeginWith:   []string{"http://", "https://"},
		EndWith:     []string{`"`},
		MustContain: []string{"blog.jobbole.com"},
	}
	re := s.GetRegexp()

	WannerFromHtml := func(html string) string {
		str := re.FindAllString(html, -1)
		ans := ""
		for _, val := range str {
			if redis.Exist([]byte(val)) {
				continue
			}
			redis.Set([]byte(val), []byte(tool.UrlKey))
			if strings.Contains(val, "comments") {
				continue
			}
			title := services.TitleFromUrl(val[:len(val)-1], ReTitle)
			if title == "" {
				continue
			}

			ans += services.MarkDownFromTitleAndURL(title, val[:len(val)-1])
			ans += "\n\n"
		}
		return ans
	}

	GlobalTelegramBot.StartSpider(URL, WannerFromHtml, dis, chatID)
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
		ans := ""
		for _, val := range str {
			if redis.Exist([]byte(val)) {
				continue
			}
			redis.Set([]byte(val), []byte(tool.UrlKey))
			title := services.TitleFromUrl(val, ReTitle)
			if title == "" {
				continue
			}

			ans += services.MarkDownFromTitleAndURL(title, val)
			ans += "\n\n"
		}
		return ans
	}
	GlobalTelegramBot.StartSpider(URL, WannerFromHtml, dis, chatID)
}
