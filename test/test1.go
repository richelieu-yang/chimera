package main

import (
	"github.com/richelieu42/chimera/v2/src/core/ioKit"
	"github.com/richelieu42/chimera/v2/src/log/logrusKit"
	"github.com/sirupsen/logrus"
)

func main() {
	logger, err := logrusKit.NewLumberjackLogger([]ioKit.LumberjackOption{ioKit.WithFilePath("a.log"), ioKit.WithConsole(true)})
	if err != nil {
		logrus.Fatal(err)
	}
	logger.Info(666)
}
