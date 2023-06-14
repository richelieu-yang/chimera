package main

import (
	"github.com/richelieu-yang/chimera/v2/src/core/ioKit"
	"github.com/richelieu-yang/chimera/v2/src/dataSizeKit"
	"github.com/richelieu-yang/chimera/v2/src/log/logrusKit"
	"github.com/sirupsen/logrus"
)

func main() {
	logrusKit.MustSetUp(nil)

	wc, err := ioKit.NewRotatableWriteCloser("test3.log", dataSizeKit.MiB*20)
	if err != nil {
		logrus.Fatal(err)
	}
	for {
		n, err := wc.Write([]byte("~!@#$%^&*()_+"))
		if err != nil {
			logrus.Fatal(err)
		}
		logrus.Info(n)
	}
}
