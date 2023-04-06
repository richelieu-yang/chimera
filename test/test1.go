package main

import (
	"fmt"
	"github.com/richelieu42/chimera/v2/src/core/floatKit"
)

func main() {
	fmt.Printf("%s\n", floatKit.ToReadableString(2.24))
	fmt.Printf("%s\n", floatKit.ToReadableString(2.0000))
	fmt.Printf("%s\n", floatKit.ToReadableString(2.000010000))
}
