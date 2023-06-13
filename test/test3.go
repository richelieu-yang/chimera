package main

import (
	"github.com/richelieu-yang/chimera/v2/src/cronKit"
	"github.com/richelieu-yang/chimera/v2/src/log/logrusKit"
	"github.com/sirupsen/logrus"
)

func main() {
	logrusKit.MustSetUp(nil)

	cron, _, err := cronKit.NewCronWithTask("@every 10s", func() {
		logrus.Info("do")
	})
	if err != nil {
		panic(err)
	}
	cron.Start()

	logrus.Info("--------")

	select {}
}
