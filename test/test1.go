package main

import (
	"fmt"
	"math"
)

func main() {
	x := math.Inf(1)

	switch {
	case x < 0, x > 0:
		fmt.Println(x) // +Inf
	case x == 0:
		fmt.Println("zero")
	default:
		fmt.Println("something else")
	}
}
