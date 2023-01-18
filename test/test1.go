package main

import (
	"fmt"
	"github.com/richelieu42/go-scales/src/core/boolKit"
)

func main() {
	v := boolKit.ToBool(nil)
	fmt.Println(v)
	v1, err := boolKit.ToBoolE(nil)
	fmt.Println(v1, err)
}
