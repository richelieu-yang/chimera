package main

import "fmt"

func main() {
	s := make([]int, 0, 5)
	fmt.Println(s[:4]) // [0 0 0 0]
	fmt.Println(s[:5]) // [0 0 0 0 0]
	fmt.Println(s[0:0])
	fmt.Println(s[0:0] != nil)
	fmt.Println(s[0:0])
	fmt.Println(s[5:5] != nil)
}
