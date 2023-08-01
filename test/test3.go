package main

import (
	"fmt"
	"github.com/richelieu-yang/chimera/v2/src/cmdKit"
	"github.com/richelieu-yang/chimera/v2/src/core/strKit"
)

func main() {
	str, err := cmdKit.ExecuteToString("sh", "-c", "ps auxw | wc -l")
	if err != nil {
		panic(err)
	}
	str = strKit.TrimSpace(str)
	fmt.Println(str)
}
