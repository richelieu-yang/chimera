package main

import (
	"github.com/panjf2000/ants/v2"
	"github.com/richelieu42/chimera/v2/src/core/ioKit"
	"github.com/richelieu42/chimera/v2/src/log/logrusKit"
	"github.com/sirupsen/logrus"
)

func main() {
	ants.Release

	logger, err := logrusKit.NewLumberjackLogger([]ioKit.LumberjackOption{
		ioKit.WithFilePath("a.log"),
		ioKit.WithConsole(true),
		//ioKit.WithMaxSize(1),
		ioKit.WithCompress(true),
	})
	if err != nil {
		logrus.Fatal(err)
	}
	for {
		logger.Info(666)
	}
}
