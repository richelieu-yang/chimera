package main

import (
	"fmt"
	"github.com/richelieu-yang/chimera/v2/src/ip/ipKit"
	"github.com/richelieu-yang/chimera/v2/src/json/jsonKit"
)

func main() {
	info, err := ipKit.GetPublicIpInfo()
	if err != nil {
		panic(err)
	}
	fmt.Println(jsonKit.MarshalIndentToString(info, "", "    "))

	//fmt.Println(ipKit.GetInternalIp()) // 172.20.10.4
	//fmt.Println()
	//fmt.Println(ipKit.GetIps()) // [172.20.10.4 198.18.0.1]
}
