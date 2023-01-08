package main

import (
	"fmt"
)

type Bean struct {
	Id int
}

func main() {
	s := []string(nil)

	fmt.Println(s[0:0])        // []
	fmt.Println(s[0:0] == nil) // true
}
