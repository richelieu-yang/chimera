package main

import (
	"github.com/richelieu42/go-scales/src/log/logrusKit"
	"github.com/richelieu42/go-scales/src/mq/rocketmq5Kit"
	"github.com/sirupsen/logrus"
)

func main() {
	logrusKit.InitializeByDefault()

	err := rocketmq5Kit.TestEndpoint("localhost:8081", "test")
	logrus.Errorf("%+v", err)
}
