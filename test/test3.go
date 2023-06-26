package main

import (
	"fmt"
	"github.com/richelieu-yang/chimera/v2/src/core/boolKit"
	"github.com/richelieu-yang/chimera/v2/src/netKit"
)

func main() {
	fmt.Println(boolKit.StringToBool("错"))

	fmt.Printf("%v\n", test(-1))
}

func test(port int) error {
	if err := netKit.AssertValidPort(port); err != nil {
		return err
	}
	return nil
}
