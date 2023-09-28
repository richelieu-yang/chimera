package main

import (
	"github.com/richelieu-yang/chimera/v2/src/log/logrusKit"
)

func main() {
	logger0 := logrusKit.NewLogger()
	logger1 := logrusKit.NewLogger()

	logger0.Info("0")
	logger1.Info("1")

	logrusKit.DisableQuote(logger1)

	logger0.Info("0")
	logger1.Info("1")
}
