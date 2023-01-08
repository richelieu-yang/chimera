package main

import "fmt"

func main() {
	s := make([]int, 0, 10)

	fmt.Println(len(s)) // 0
	fmt.Println(cap(s)) // 10
	fmt.Println(s[0])
}
