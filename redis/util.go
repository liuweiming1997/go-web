package redis

import "github.com/sirupsen/logrus"

func showErr(method string, err error) {
	logrus.Errorf("[redis] doing %s error %v\n", method, err)
}
