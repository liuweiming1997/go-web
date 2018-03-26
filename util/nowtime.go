package util

import "time"

func NowTime() string {
	return time.Now().String()[0:19]
}

func NowDate() string {
	return time.Now().String()[0:10]
}
