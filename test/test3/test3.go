package main

import (
	"github.com/richelieu-yang/chimera/v2/src/validateKit"
)

func main() {
	err := validateKit.Var("", "hostname_port")
	if err != nil {
		panic(err)
	}
}
