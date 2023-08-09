package main

import (
	"github.com/richelieu-yang/chimera/v2/src/log/logrusKit"
	"github.com/richelieu-yang/chimera/v2/src/statKit"
)

func main() {
	logrusKit.MustSetUp(&logrusKit.Config{
		Level:      "",
		PrintBasic: true,
	})

	statKit.MustSetup("")
	select {}

	//statKit.PrintStats(nil)
}
