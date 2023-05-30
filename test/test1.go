package main

import "fmt"

func main() {
	type bean struct {
	}

	var b *bean
	fmt.Println(b == nil) // true
	fmt.Println(&b)       // 0x1400000e028
}
