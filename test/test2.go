package main

import (
	"fmt"
)

func main() {
	defer func() {
		fmt.Println("a")
	}()

	{
		defer func() {
			fmt.Println("b")
		}()
	}

	defer func() {
		fmt.Println("c")
	}()
}
