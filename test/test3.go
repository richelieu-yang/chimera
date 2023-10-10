package main

import (
	"fmt"
	"github.com/rs/xid"
)

func main() {
	for i := 0; i < 1000; i++ {
		str := xid.New().String()
		fmt.Println(str)
		fmt.Println(len(str))
		if len(str) != 20 {
			panic(len(str))
		}
	}
}
