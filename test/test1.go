package main

import (
	"fmt"
	"github.com/richelieu42/go-scales/src/core/sliceKit"
)

func main() {
	s := sliceKit.Merge([]string(nil), []string{"a", "b"}, []string(nil), []string{"b", "c"})
	fmt.Println(s)
}
