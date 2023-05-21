package main

import (
	"fmt"
)

func main() {
	s1 := make([]int, 0, 5)
	fmt.Println("s1切片: ", s1, len(s1), cap(s1))

	appendFunc1(s1)

	fmt.Println("s1切片: ", s1, len(s1), cap(s1))
	fmt.Println("s1切片表达式: ", s1[:5])
}

func appendFunc1(s2 []int) {
	//s2 = append(s2, 1, 2, 3)
	s2 = append(s2, 1, 2, 3, 4, 5, 6, 7, 8, 9)
	fmt.Println("s2切片: ", s2, len(s2), cap(s2))
}
