package main

import (
	"github.com/sirupsen/logrus"
	"time"
)

func main() {
	timer := time.AfterFunc(time.Second*3, func() {
		logrus.Info("time up")
	})
	time.Sleep(time.Second)
	logrus.Info(timer.Stop()) // true
	logrus.Info(timer.Stop()) // false

	time.Sleep(time.Second * 6)
}
