package main

import (
	"github.com/richelieu42/chimera/v2/src/core/signalKit"
	"github.com/richelieu42/chimera/v2/src/log/logrusKit"
	"os"
	"time"
)

func main() {
	logrusKit.MustSetUp(nil)
	signalKit.MonitorExitSignal(func(signal os.Signal) {
		time.Sleep(time.Second * 3)
	})

	select {}
}
