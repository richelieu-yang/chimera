package main

import "fmt"

func main() {
	// 支持遍历整数
	for i := range 3 {
		fmt.Println(i)
	}
	fmt.Println("go1.22 has lift-off!")
}
