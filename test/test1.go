package main

import "github.com/richelieu42/chimera/src/log/logrusKit"

func main() {
	logrusKit.SetUp(&logrusKit.Config{
		PrintBasic: true,
	})
}
