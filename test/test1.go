package main

import (
	"github.com/richelieu42/chimera/src/copyKit"
	"github.com/richelieu42/chimera/src/log/logrusKit"
	"github.com/sirupsen/logrus"
	"unsafe"
)

func main() {
	var formatter = logrusKit.NewTextFormatter("")

	formatter1, err := copyKit.DeepCopy(formatter)
	if err != nil {
		logrus.Fatal(err)
	}

	formatter.ForceQuote = false
	formatter.DisableQuote = false
	formatter1.ForceQuote = true
	formatter1.DisableQuote = true

	logrus.Info(unsafe.Pointer(&formatter.FieldMap))
	logrus.Info(unsafe.Pointer(&formatter1.FieldMap))
}
