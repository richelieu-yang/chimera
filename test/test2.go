package main

import "fmt"

func main() {
	var s []string = nil

	for i := 0; i < len(s); i++ {
		fmt.Println(i, s[i])
	}
	fmt.Println("--------------------")
}
