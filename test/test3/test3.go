package main

import (
	"fmt"
	"github.com/richelieu-yang/chimera/v2/src/core/signalKit"
	"github.com/richelieu-yang/chimera/v2/src/processKit"
	"os"
)

func main() {
	signalKit.MonitorExitSignals(func(sig os.Signal) {
		fmt.Println("0", sig.String())
	})
	signalKit.MonitorExitSignals(func(sig os.Signal) {
		fmt.Println("1", sig.String())
	})

	fmt.Println("pid:", processKit.PID)

	select {}
}
