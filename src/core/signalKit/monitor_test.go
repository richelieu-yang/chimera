package signalKit

import (
	"fmt"
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

	select {}
}
