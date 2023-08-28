package main

import (
	"fmt"
	"github.com/richelieu-yang/chimera/v2/src/core/timeKit"
	"time"
)

func main() {
	fmt.Println(timeKit.GetNetworkTime())

	time.Sleep(time.Second * 10)
}
