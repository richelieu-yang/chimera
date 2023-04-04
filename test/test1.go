package main

import (
	"fmt"
	"github.com/richelieu42/chimera/src/idKit"
)

func main() {
	fmt.Println(idKit.NewUUID())
	fmt.Println(idKit.NewSimpleUUID())
}
