package signalKit

import (
	"fmt"
	"github.com/richelieu-yang/chimera/v3/src/processKit"
	"os"
	"testing"
)

func TestMonitorExitSignals(t *testing.T) {
	MonitorExitSignals(func(sig os.Signal) {
		fmt.Println("0", sig.String())
	})
	MonitorExitSignals(func(sig os.Signal) {
		fmt.Println("1", sig.String())
	})

	fmt.Println(processKit.PID)

	select {}
}
