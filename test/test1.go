package main

import (
	"github.com/sirupsen/logrus"
	"time"
)

func main() {
	ch := make(chan int)

	logrus.Info("------")

	go func() {
		time.Sleep(time.Second)
		logrus.Info("[goroutine] close starts")
		close(ch)
		logrus.Info("[goroutine] close ends")
	}()

	logrus.Info("write starts")
	ch <- 1 // panic: send on closed channel
	logrus.Info("write ends")
}
