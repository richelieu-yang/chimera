package main

import (
	"fmt"
	"github.com/richelieu-yang/chimera/v2/src/core/strKit"
	"github.com/richelieu-yang/chimera/v2/src/core/timeKit"
)

func main() {
	fmt.Sprintf("%d %s", 1, "a")
	strKit.Format("%d %s", 1, "a")

	fmt.Println(timeKit.GetNetworkTime())
}
