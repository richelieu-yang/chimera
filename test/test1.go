package main

import (
	"fmt"
	"github.com/richelieu42/chimera/v2/src/assertKit"
)

func main() {
	fmt.Println(assertKit.AssertHttpUrl("https:/github.com/go-playground/validator"))
}
