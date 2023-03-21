package main

import (
	"fmt"
	"github.com/richelieu42/chimera/src/core/intKit"
)

func main() {
	fmt.Println(intKit.ToIntE(nil))
	fmt.Println(intKit.ToIntE(false))
	fmt.Println(intKit.ToIntE(true))
	fmt.Println(intKit.ToIntE(""))
}
