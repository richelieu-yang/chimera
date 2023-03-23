package main

import (
	"fmt"
	"github.com/richelieu42/chimera/src/core/sliceKit"
)

func main() {
	flag := sliceKit.IsSortedByKey([]string{"a", "aa", "bb", "ccc"}, func(s string) int {
		return len(s)
	})
	fmt.Println(flag) // true
}
