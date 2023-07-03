package main

import (
	"fmt"
)

type SampleStruct struct {
	Value string
}

func main() {
	var keys = make([]string, 0, 60)
	fmt.Println(len(keys))
	fmt.Println(cap(keys))
}
