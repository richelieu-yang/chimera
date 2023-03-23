package main

import (
	"fmt"
	"github.com/richelieu42/chimera/src/core/sliceKit"
	"strings"
)

func main() {
	s := sliceKit.FilterAndRevise([]string{"cpu", "gpu", "mouse", "keyboard"}, func(item string, index int) (string, bool) {
		if strings.HasSuffix(item, "pu") {
			return "right-" + item, true
		}
		return "", false
	})
	fmt.Println(s) // [right-cpu right-gpu]
}
