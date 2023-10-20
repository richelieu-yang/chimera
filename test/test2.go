package main

import (
	"fmt"
	"github.com/richelieu-yang/chimera/v2/src/netKit"
)

func main() {
	fmt.Println(netKit.JoinHostnameAndPort("127.0.0.1", 80)) // 127.0.0.1:80
	fmt.Println(netKit.JoinHostnameAndPort("", 8888))        // :8888
}
