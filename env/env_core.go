package env

import "github.com/sundayfun/go-web/tool"
import "os"

func GetTelegramToken() string {
	return os.Getenv(tool.TelegramToken)
}

func GetRedisHost() string {
	return os.Getenv(tool.RedisHost)
}
