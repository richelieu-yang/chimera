package main

import (
	"fmt"
	"github.com/richelieu-yang/chimera/v2/src/validateKit"
)

func main() {
	cap := 2000 - 1
	err := validateKit.Var(cap, "gte=2000")
	fmt.Println(err)
}
