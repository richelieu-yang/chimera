package main

import "github.com/richelieu42/chimera/v2/src/core/osKit"

func main() {
	osKit.MonitorExitSignal(nil)

	select {}
}
