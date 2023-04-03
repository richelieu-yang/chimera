package main

import (
	"fmt"
)

func main() {
	var s []int = []int{}
	s1 := s[0:0]

	fmt.Println(s1)
	fmt.Println(s1 != nil)

	//s1 := sliceKit.Intercept(s, len(s), len(s))
	//fmt.Println(s1)        // []
	//fmt.Println(s1 != nil) // true
}
