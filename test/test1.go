package main

import (
	"github.com/richelieu-yang/chimera/v2/src/core/cpuKit"
	"github.com/richelieu-yang/chimera/v2/src/log/logrusKit"
	"github.com/sirupsen/logrus"
	"time"
)

func init() {
	cpuKit.SetUp()
}

func main() {
	logrusKit.MustSetUp(nil)

	logrus.Info("sleep starts")
	time.Sleep(time.Second * 3)
	logrus.Info("sleep ends")
}
