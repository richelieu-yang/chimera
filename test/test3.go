package main

import (
	"fmt"
	"github.com/richelieu-yang/chimera/v2/src/log/logrusKit"
	"github.com/richelieu-yang/chimera/v2/src/processKit"
	"github.com/sirupsen/logrus"
)

func init() {
	logrusKit.MustSetUp(nil)
}

func main() {
	defer func() {
		fmt.Println("***")
		logrus.Info("***")
		//time.Sleep(time.Second * 10)
		logrus.Info("===")
	}()

	fmt.Println("PID", processKit.PID)

	for {
	}
}
