package main

import (
	"fmt"
	"github.com/richelieu-yang/chimera/v2/src/cmd/cmdKit"
	"github.com/richelieu-yang/chimera/v2/src/log/logrusKit"
)

func init() {
	logrusKit.MustSetUp(nil)
}

func main() {
	fmt.Println(cmdKit.LookPath("pip"))
}
