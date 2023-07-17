package main

import (
	"github.com/richelieu-yang/chimera/v2/src/cronKit"
	"github.com/sirupsen/logrus"
)

func main() {
	c, _, err := cronKit.NewCronWithTask("0 0 0 * * *", func() {
		logrus.Info("execute")
	})
	if err != nil {
		panic(err)
	}
	c.Run()
}
