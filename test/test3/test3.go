package main

import (
	"fmt"
	"github.com/richelieu-yang/chimera/v2/src/validateKit"
)

func main() {
	fmt.Println(validateKit.Var("127.0.0.255", "omitempty,hostname|ipv4"))
}
