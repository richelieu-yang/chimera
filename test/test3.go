package main

import (
	"github.com/richelieu42/chimera/v2/src/core/timeKit"
	"github.com/richelieu42/chimera/v2/src/log/logrusKit"
	"github.com/sirupsen/logrus"
	"runtime"
	"time"
)

func main() {
	logrusKit.MustSetUp(nil)

	logrus.Info("===", runtime.NumGoroutine())
	t, url, err := timeKit.GetNetworkTime()
	if err != nil {
		logrus.Fatal(err)
	}
	logrus.Info(t, url)
	logrus.Info("===", runtime.NumGoroutine())
	time.Sleep(time.Second * 10)
	logrus.Info("===", runtime.NumGoroutine())
}
