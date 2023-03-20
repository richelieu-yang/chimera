package main

import (
	"fmt"
	"github.com/richelieu42/chimera/src/core/ioKit"
	"github.com/sirupsen/logrus"
	"time"
)

func main() {
	wc, err := ioKit.NewRotateFileWriteCloser("/Users/richelieu/Downloads/c.log", time.Hour, time.Hour*12, true)
	//wc, err := ioKit.NewDailyRotateRuleWriteCloser("/Users/richelieu/Downloads/c.log", "-", 1, false)
	if err != nil {
		fmt.Println(err.Error())
		logrus.Fatal(err)
	}
	if _, err := wc.Write([]byte("123\n")); err != nil {
		logrus.Fatal(err)
	}
	if _, err := wc.Write([]byte("456\n")); err != nil {
		logrus.Fatal(err)
	}
	logrus.Info("------------------")

	logger := logrus.New()
	logger.Out = wc
	logger.Info("789")

	//_ = wc.Close()

	//time.Sleep(time.Second * 3)
}
