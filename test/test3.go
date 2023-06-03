package main

import (
	"fmt"
	"os"
)

func main() {
	err := os.MkdirAll(" ", os.ModePerm)
	fmt.Println(err)
	fmt.Println(err == nil)
}
