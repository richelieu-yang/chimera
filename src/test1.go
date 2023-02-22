package main

import (
	"fmt"
	"time"
)

func main() {
	go func() {
		time.Sleep(time.Second)
		panic("p")
	}()
	time.Sleep(time.Second * 3)
	fmt.Println("================")
}
