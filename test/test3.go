package main

import (
	"fmt"
	"github.com/richelieu-yang/chimera/v2/src/component/web/push/pushKit"
)

func main() {
	var c pushKit.Channel = nil

	c1, ok := c.(pushKit.Channel)
	fmt.Println(c1, ok)
}
