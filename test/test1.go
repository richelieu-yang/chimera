package main

import (
	"fmt"
	"github.com/richelieu42/go-scales/src/core/strKit"
)

func main() {
	str := strKit.ToString(nil)
	fmt.Println(str)
	str1, err := strKit.ToStringE(nil)
	fmt.Println(str1, err)
}
