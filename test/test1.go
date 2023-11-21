package main

import (
	"github.com/sirupsen/logrus"
	"time"
)

func main() {
	i := NewInterval(func(t time.Time) {
		logrus.Info(t)
	}, time.Second*3)

	logrus.Info("sleep starts")
	time.Sleep(time.Second * 10)
	logrus.Info("sleep ends")

	i.Stop()
	i.Stop()
	i.Stop()

	logrus.Info("sleep starts")
	time.Sleep(time.Second * 10)
	logrus.Info("sleep ends")
}
