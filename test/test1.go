package main

import "fmt"

func main() {
	var s []int = nil
	fmt.Println(s[:] == nil)
	fmt.Println(s[:0] == nil)
	fmt.Println(s[0:] == nil)
	fmt.Println(s[0:0] == nil)
}
