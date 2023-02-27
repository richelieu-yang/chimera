package main

import (
	"fmt"
	"github.com/richelieu42/go-scales/src/core/sliceKit"
)

func main() {
	old := []string{"0", "1", "2"}
	texts := sliceKit.Copy2(old)

	texts[0] = "3"

	fmt.Println(old)
	fmt.Println(texts)

	//texts1, _ := sliceKit.Remove(texts, "1")
	//fmt.Println(old)    // [0 2 2]
	//fmt.Println(texts)  // [0 2 2]
	//fmt.Println(texts1) // [0 2]
}
