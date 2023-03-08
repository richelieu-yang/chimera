package main

import (
	"fmt"
)

type MyInt int

type Bean struct {
}

func main() {
	fmt.Println(1)       // 1024
	fmt.Println(1 << 10) // 1024
	fmt.Println(1 << 20) // 1048576
	fmt.Println(1 << 30) // 1073741824
	fmt.Println(1 << 40) // 1099511627776
}
