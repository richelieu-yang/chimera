package main

import (
	"fmt"
	"github.com/richelieu42/go-scales/src/core/sliceKit"
)

func main() {
	m := sliceKit.SliceToMap[string, string, string]([]string{"0", "1"}, nil)
	fmt.Println(m) // map[key0:value0 key1:value1]
}
