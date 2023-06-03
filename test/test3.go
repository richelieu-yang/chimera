package main

import (
	"github.com/richelieu42/chimera/v2/src/core/fileKit"
	"github.com/richelieu42/chimera/v2/src/log/logrusKit"
	"github.com/sirupsen/logrus"
)

func main() {
	logrusKit.MustSetUp(nil)

	err2 := a()
	logrusKit.DisableQuote(nil)
	logrus.Fatalf("%+v", err2)
}

func a() error {
	return fileKit.AssertExist("")
}
