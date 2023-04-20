package main

import (
	"fmt"
	"github.com/richelieu42/chimera/v2/src/assertKit"
)

func main() {
	fmt.Println(assertKit.AssertIPv4("192.168.9.254"))
	fmt.Println(assertKit.AssertIPv4("127.0.0.1"))
	fmt.Println(assertKit.AssertIPv4("localhost"))
	fmt.Println(assertKit.AssertIPv4("::1"))
}
