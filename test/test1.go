package main

import (
	"fmt"
	"unsafe"
)

func main() {
	s := []int{0, 1, 2, 3}
	s1 := s[1:]

	fmt.Println(s, unsafe.Pointer(&s))
	fmt.Println(s1, unsafe.Pointer(&s1))

	s1[2] = 9
	fmt.Println(s, unsafe.Pointer(&s))
	fmt.Println(s1, unsafe.Pointer(&s1))
}
