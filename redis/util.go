package redis

import "github.com/sirupsen/logrus"

func showErr(method string, err error) {
	logrus.Errorf("doing %s error %v\n", method, err)
}
