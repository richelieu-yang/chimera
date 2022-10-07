package main

import (
	"fmt"
	"unsafe"
)

func main() {
	str := "abc"
	a := len(str)
	b := unsafe.Sizeof(str)
	fmt.Println(a, b) // 3 16
}
