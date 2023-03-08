package main

import (
	"fmt"
)

type MyInt int

type Bean struct {
}

func main() {
	fmt.Println(1 << 10)
	fmt.Println(1 << 20)
}
