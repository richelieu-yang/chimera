package main

import (
	"fmt"
	"github.com/richelieu-yang/chimera/v2/src/ip/ipRegionKit"
)

func main() {
	ipRegionKit.MustSetUp("_chimera-lib/ip2region.xdb")
	fmt.Println(ipRegionKit.GetRegion("1.0.16.0"))
}
