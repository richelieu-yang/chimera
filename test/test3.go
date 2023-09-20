package main

import "github.com/richelieu-yang/chimera/v2/src/ip/ipKit"

func main() {
	ip := "127.001"

	if err := ipKit.AssertIPv4(ip); err != nil {
		panic(err)
	}
}
