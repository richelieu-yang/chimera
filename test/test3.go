package main

import (
	"flag"
	"fmt"
)

var addr string

func init() {
	flag.StringVar(&addr, "addr", "default", "address")
}

func main() {
	// 必要，否则取不到值
	flag.Parse()

	fmt.Printf("addr:[%s]\n", addr)
}
