package main

import (
	"github.com/richelieu42/go-scales/src/log/logrusKit"
	"github.com/sirupsen/logrus"
	"time"
)

func main() {
	logger, err := logrusKit.NewRotateFileLogger("/Users/richelieu/Downloads/111.log", nil, logrus.DebugLevel, time.Second*10, time.Second*30, false)
	if err != nil {
		panic(err)
	}
	logger.Info("666")
	logger.Info("999")
	logrus.Info("-------------------")
}
