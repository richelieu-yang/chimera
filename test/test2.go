package main

import (
	"fmt"
)

func main() {
LABEL1:
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			fmt.Printf("i = [%d], j = [%d].\n", i, j)
			if j == 1 {
				fmt.Println("goto")
				goto LABEL1
			}
		}
	}
}
