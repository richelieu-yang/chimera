package main

import (
	"fmt"
	"github.com/richelieu42/chimera/v2/src/idKit"
)

func main() {
	str := idKit.NewUUID()
	fmt.Println(len(str), str)

	str = idKit.NewSimpleUUID()
	fmt.Println(len(str), str)
}
