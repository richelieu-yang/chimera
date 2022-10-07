package main

import (
	"fmt"
)

func main() {
	fmt.Println("---------------------------------------------")
	fmt.Println(t0())
	fmt.Println("---------------------------------------------")
}

func t0() (str string) {
	str = "default"

	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
			str = "0"
		} else {
			str = "-1"
		}
	}()

	foo()
	str = "1"
	return
}

func foo() {
	panic("test panic")
}
