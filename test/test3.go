package main

import (
	"fmt"
	"github.com/duke-git/lancet/v2/netutil"
)

func main() {
	fmt.Println(netutil.GetInternalIp())

	fmt.Println(netutil.GetRequestPublicIp())
}
