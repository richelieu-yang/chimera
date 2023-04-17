package main

import (
	"github.com/richelieu42/chimera/v2/src/core/runtimeKit"
	"github.com/richelieu42/chimera/v2/src/core/signalKit"
	"github.com/richelieu42/chimera/v2/src/log/logrusKit"
	"github.com/sirupsen/logrus"
	"os"
	"time"
)

func main() {
	//logrus.Info(syscall.SIGHUP.String())
	//logrus.Info(syscall.SIGINT.String())

	logrus.Infof("pid: [%d]", runtimeKit.PID)

	logrusKit.MustSetUp(nil)
	signalKit.MonitorExitSignal(func(signal os.Signal) {
		time.Sleep(time.Second * 3)
	})

	select {}
}
