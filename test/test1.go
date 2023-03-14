package main

import (
	"github.com/sirupsen/logrus"
	"time"
)

func main() {
	ch := make(chan struct{})
	go func() {
		time.Sleep(1 * time.Second)
		close(ch)
	}()

	logrus.Info("脑子好像进...")
	<-ch
	logrus.Info("煎鱼了！")
}
