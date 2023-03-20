package main

import (
	"fmt"
	"github.com/richelieu42/chimera/src/core/ioKit"
	"github.com/sirupsen/logrus"
)

func main() {
	wc, err := ioKit.NewDailyRotateRuleWriteCloser("a.log", "-", 1, false)
	//wc, err := ioKit.NewSizeLimitRotateRuleWriteCloser("/Users/richelieu/Downloads", "-", 1, 10, 10, false)
	if err != nil {
		fmt.Println(err.Error())
		logrus.Fatal(err)
	}
	if _, err := wc.Write([]byte("123")); err != nil {
		logrus.Fatal(err)
	}
	if _, err := wc.Write([]byte("456")); err != nil {
		logrus.Fatal(err)
	}
	logrus.Info("------------------")

	wc.Close()
}
