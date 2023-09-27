package main

import (
	"github.com/richelieu-yang/chimera/v2/src/validateKit"
)

func main() {
	if err := validateKit.Port(65535); err != nil {
		panic(err)
	}
}
