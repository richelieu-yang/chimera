package main

import (
	"fmt"
	"math/rand"
	"time"
)

type MyInt int

type Bean struct {
}

func main() {
	rand := rand.New(rand.NewSource(time.Now().UnixNano()))

	rand.Intn()

	fmt.Println(1)       // 1024
	fmt.Println(1 << 10) // 1024
	fmt.Println(1 << 20) // 1048576
	fmt.Println(1 << 30) // 1073741824
	fmt.Println(1 << 40) // 1099511627776
}
