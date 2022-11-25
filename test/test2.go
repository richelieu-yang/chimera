package main

import (
	"github.com/richelieu42/go-scales/src/log/logrusKit"
	"github.com/sirupsen/logrus"
	"time"
)

func main() {
	//logger, err := logrusKit.NewFileLogger("/Users/richelieu/Downloads/a.txt", nil, logrus.DebugLevel, false)
	logger, err := logrusKit.NewRotateFileLogger("/Users/richelieu/Downloads/a.txt", nil, logrus.DebugLevel, time.Second*5, time.Second*30, false)
	if err != nil {
		panic(err)
	}

	logger.Info(0)
	if err := logrusKit.DisposeLogger(logger); err != nil {
		panic(err)
	}
	logger.Info(1)

	//for i := 0; i < 60; i++ {
	//	logger.Info(i)
	//	time.Sleep(time.Second)
	//}
	//logger.Info("666")
	//if err := logrusKit.DisposeLogger(logger); err != nil {
	//	panic(err)
	//}
	//logger.Info("777")

}
