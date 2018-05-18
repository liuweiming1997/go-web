package tool

import "github.com/sirupsen/logrus"

func CheckErr(err error) {
	if err != nil {
		logrus.Error(err)
	}
}
