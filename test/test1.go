package main

import (
	"fmt"
	"github.com/richelieu-yang/chimera/v2/src/core/errorKit"
)

func main() {
	fmt.Println(a())
}

func a() error {
	return errorKit.New("zzz")
}
