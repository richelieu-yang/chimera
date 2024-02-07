package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	fmt.Println(runtime.Version())

	values := []int{1, 2, 3, 4, 5}

	for _, val := range values {
		go func() {
			fmt.Printf("%d\n", val)
		}()
	}
	time.Sleep(time.Second * 3)
}
