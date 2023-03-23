package main

import (
	"fmt"
	"github.com/richelieu42/chimera/src/core/sliceKit"
)

func main() {
	s := []string{"", "foo", "", "bar", ""}
	s1 := sliceKit.Compact[string](s)
	fmt.Println(s1) // [foo bar]
}
