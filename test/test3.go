package main

import (
	"fmt"
	"github.com/richelieu-yang/chimera/v2/src/core/timeKit"
	"time"
)

func main() {
	now := time.Now()
	fmt.Println(now) // 2023-08-18 15:24:03.167655 +0800 CST m=+0.004041126
	t := timeKit.ToZeroAM(now)
	fmt.Println(t) // 2023-08-18 00:00:00 +0800 CST
}
