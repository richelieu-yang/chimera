package main

import (
	"github.com/richelieu-yang/chimera/v2/src/log/logrusKit"
	"github.com/richelieu-yang/chimera/v2/src/statKit"
)

func main() {
	logrusKit.MustSetUp(nil)

	statKit.MustSetup("")
	select {}
}
