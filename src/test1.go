package main

import (
	"fmt"
	"github.com/richelieu42/go-scales/src/core/sliceKit"
)

func main() {
	texts := []string{"0", "1", "2"}
	fmt.Println(texts) // [0 1 2]

	texts1, _ := sliceKit.Remove(texts, "1")
	fmt.Println(texts)  // [0 2 2]
	fmt.Println(texts1) // [0 2]
}
