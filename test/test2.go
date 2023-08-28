package main

import (
	"fmt"
	"github.com/richelieu-yang/chimera/v2/src/core/timeKit"
	"github.com/richelieu-yang/chimera/v2/src/log/logrusKit"
	"time"
)

func main() {
	logrusKit.MustSetUp(nil)

	fmt.Println(timeKit.GetNetworkTime())

	time.Sleep(time.Second * 10)
}
