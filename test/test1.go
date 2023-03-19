package main

import (
	"fmt"
	"github.com/richelieu42/go-scales/src/core/sliceKit"
)

func main() {
	s := sliceKit.Uniq([]interface{}{0, 1, 2, 0, "1", "2", "1"})
	fmt.Println(s)
}
