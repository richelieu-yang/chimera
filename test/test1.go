package main

import (
	"fmt"
	"github.com/richelieu42/chimera/v2/src/randomKit"
)

func main() {
	for i := 0; i < 10; i++ {
		fmt.Println(randomKit.Int(-2, 2))
	}
}
