package signalKit

import (
	"fmt"
	"github.com/richelieu-yang/chimera/v3/src/processKit"
	"github.com/sirupsen/logrus"
	"os"
	"testing"
	"time"
)

func TestMonitorExitSignals(t *testing.T) {
	MonitorExitSignals(func(sig os.Signal) {
		logrus.Info("0", sig.String())
		time.Sleep(time.Second * 3)
	})
	MonitorExitSignals(func(sig os.Signal) {
		logrus.Info("1", sig.String())
		time.Sleep(time.Second * 3)
	})

	fmt.Println(processKit.PID)

	select {}
}
