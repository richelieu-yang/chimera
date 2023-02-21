package main

import (
	"github.com/richelieu42/go-scales/src/log/logrusKit"
	"github.com/sirupsen/logrus"
	"time"
)

func main() {
	logger, err := logrusKit.NewRotateFileLogger("/Users/richelieu/Downloads/c.txt", nil, logrus.DebugLevel, time.Hour, 12*time.Hour, true, true)
	if err != nil {
		panic(err)
	}
	logger.Debug("123")
	if err := logrusKit.DisposeLogger(logger); err != nil {
		panic(err)
	}
	if err := logrusKit.DisposeLogger(logger); err != nil {
		panic(err)
	}
	logger.Debugf("%d", 456)
	logger.Debugf("%d", 789)
}
