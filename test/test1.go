package main

import (
	"fmt"
	"github.com/richelieu42/chimera/src/core/sliceKit"
)

func main() {
	s := []string{"a", "aa", "aaa", "aa", "aa"}
	s1 := sliceKit.DropRightWhile(s, func(val string) bool {
		return len(val) <= 2
	})
	fmt.Println(s)  // [a aa aaa aa aa]
	fmt.Println(s1) // [aaa aa aa]
}
