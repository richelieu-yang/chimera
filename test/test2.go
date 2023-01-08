package main

import (
	"fmt"
	"github.com/richelieu42/go-scales/src/core/sliceKit"
)

type Bean struct {
	Id int
}

func main() {
	s1, err := sliceKit.DeepCopy([]string(nil))
	if err != nil {
		panic(err)
	}
	fmt.Println(s1 == nil)
}
