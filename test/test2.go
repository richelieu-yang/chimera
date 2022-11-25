package main

import (
	"github.com/richelieu42/go-scales/src/log/logrusKit"
	"github.com/sirupsen/logrus"
)

func main() {
	logger, err := logrusKit.NewFileLogger("/Users/richelieu/Downloads/a.txt", nil, logrus.DebugLevel, false)
	if err != nil {
		panic(err)
	}

	logger.Info("666")
	if err := logrusKit.DisposeLogger(logger); err != nil {
		panic(err)
	}
	logger.Info("777")

}
