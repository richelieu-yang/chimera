package main

import (
	"fmt"
	"sync/atomic"
)

func main() {

	var i int32 = 0

	fmt.Println(i) // 0
	atomic.AddInt32(&i, -1)
	atomic.AddInt32(&i, -1)
	fmt.Println(i) // -2

	atomic.Compa

}
