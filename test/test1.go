package main

import "fmt"

func main() {
	var s []int = nil

	fmt.Println(s[:] == nil)   // true
	fmt.Println(s[0:] == nil)  // true
	fmt.Println(s[:0] == nil)  // true
	fmt.Println(s[0:0] == nil) // true
	fmt.Println(s[1:1] == nil) // panic: runtime error: slice bounds out of range [:1] with capacity 0
}
