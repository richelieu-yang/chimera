package main

import "fmt"

func main() {
	var s []interface{}
	var s1 []interface{}

	s = append(s)
	fmt.Println(s == nil) // true

	s = append(s, s1...)
	fmt.Println(s == nil) // true
}
