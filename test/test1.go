package main

import (
	"fmt"
	"github.com/richelieu42/chimera/v2/src/core/strKit"
)

func main() {
	str := "http://127.0.0.1?"
	length := len(str)

	fmt.Println(strKit.Index(str, "?"))
	fmt.Println(length)

}
