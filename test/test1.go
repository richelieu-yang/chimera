package main

import (
	"fmt"
	"github.com/richelieu42/go-scales/src/randomKit"
	"time"
)

func main() {
	for i := 0; i < 1000; i++ {
		go func() {
			fmt.Println(randomKit.Bool())
		}()
	}
	time.Sleep(time.Second * 3)
}
