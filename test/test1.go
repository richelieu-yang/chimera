package main

import (
	"fmt"
	"time"
)

func main() {
	for i := 100; i < 110; i++ {
		go test(i)
	}

	time.Sleep(time.Second * 6)
}

func test(i int) {
	if i == 100 {
		time.Sleep(time.Second * 3)
	}
	fmt.Println(i)
}
