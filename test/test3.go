package main

import (
	"github.com/richelieu-yang/chimera/v2/src/cronKit"
	"github.com/richelieu-yang/chimera/v2/src/log/logrusKit"
	"github.com/sirupsen/logrus"
)

func main() {
	logrusKit.MustSetUp(nil)

	c := cronKit.NewCron()
	_, err := c.AddFunc("0 0/2 * * * ?", func() {
		logrus.Info("-")
	})
	if err != nil {
		panic(err)
	}
	logrus.Info("-")
	c.Run()
}
