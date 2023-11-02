package main

import (
	"fmt"
	"time"
)

func main() {
	s := []string{"0", "1", "2"}

	for _, ele := range s {
		go func() {
			fmt.Println(ele)
		}()
	}
	time.Sleep(time.Second)
}
