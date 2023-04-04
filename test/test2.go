package main

import (
	"fmt"
)

type DemoError struct {
}

func (de *DemoError) Error() string {
	return "DemoError"
}

func main() {
	var err error = nil
	var obj interface{} = nil

	fmt.Println(err == obj)
}
