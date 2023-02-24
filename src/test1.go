package main

import (
	"github.com/richelieu42/go-scales/src/mq/rocketmq5Kit"
	"github.com/sirupsen/logrus"
)

func main() {
	err := rocketmq5Kit.TestEndpoint("localhost:8081", "test")
	if err != nil {
		panic(err)
	}
	logrus.Info("-------------------------------")
}
