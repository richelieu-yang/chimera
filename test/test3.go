package main

import (
	"fmt"
	"github.com/duke-git/lancet/v2/random"
)

func main() {
	for i := 0; i < 100000; i++ {
		tmp := random.RandInt(0, 100)
		if tmp == 100 {
			panic("--------")
		}
		fmt.Println(tmp)
	}

	random.UUIdV4()

}
