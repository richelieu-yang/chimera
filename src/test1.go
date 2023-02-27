package main

import (
	"github.com/richelieu42/go-scales/src/log/logrusKit"
	"github.com/richelieu42/go-scales/src/mq/rocketmq5Kit"
	"github.com/sirupsen/logrus"
)

func main() {
	logrusKit.InitializeByDefault()

	logrus.Info("start")
	err := rocketmq5Kit.VerifyEndpoint("localhost:8081", "test")
	if err != nil {
		panic(err)
	}
	logrus.Info("success")
}
