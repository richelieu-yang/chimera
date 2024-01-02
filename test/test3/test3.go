package main

import "github.com/richelieu-yang/chimera/v2/src/log/logrusKit"

func main() {
	logrusKit.MustSetUp(&logrusKit.Config{
		Level:      "",
		PrintBasic: true,
	})
}
