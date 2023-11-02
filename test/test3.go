package main

import (
	"fmt"
	"github.com/richelieu-yang/chimera/v2/src/core/sliceKit"
	"github.com/richelieu-yang/chimera/v2/src/validateKit"
)

func main() {

	sliceKit.Contains()

	cap := 2000 - 1
	err := validateKit.Var(cap, "gte=2000")
	fmt.Println(err)
}
