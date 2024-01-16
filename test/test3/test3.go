package main

import (
	"fmt"
	"github.com/richelieu-yang/chimera/v2/src/core/signalKit"
	"os"
)

func main() {
	signalKit.MonitorExitSignals(func(sig os.Signal) {
		fmt.Println("0", sig.String())
	})
	signalKit.MonitorExitSignals(func(sig os.Signal) {
		fmt.Println("1", sig.String())
	})

	select {}
}
