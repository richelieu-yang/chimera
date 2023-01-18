package main

import (
	"fmt"
	"github.com/richelieu42/go-scales/src/core/intKit"
)

func main() {
	v := intKit.ToInt(nil)
	fmt.Println(v)
	v1, err := intKit.ToIntE(nil)
	fmt.Println(v1, err)
}
