package main

import (
	"fmt"
	"github.com/richelieu42/go-scales/src/core/sliceKit"
)

func main() {
	//var s []string = nil
	//var s1 []string = nil

	//sliceKit.Append()

	s := sliceKit.Append([]string(nil), "0")
	fmt.Println(s)
	fmt.Println(s == nil)
	fmt.Println(len(s))

	s1 := []string{"0"}
	fmt.Println(s1)
}
