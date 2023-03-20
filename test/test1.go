package main

import (
	"github.com/richelieu42/chimera/src/log/logrusKit"
	"github.com/sirupsen/logrus"
)

func main() {
	logger := logrusKit.NewLogger(nil, logrus.PanicLevel)

	logger.Trace("Trace")
	logger.Debug("Debug")
	logger.Info("Info")
	logger.Warn("Warn")
	logger.Error("Error")
}
