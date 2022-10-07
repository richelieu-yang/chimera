package main

import (
	"fmt"
	"gitee.com/richelieu042/go-scales/src/core/timeKit"
	"time"
)

func main() {
	ticker := timeKit.SetInterval(func(t time.Time) {
		fmt.Println(t)
	}, time.Second*5)

	fmt.Println("--------------------------------")
	time.Sleep(time.Second * 11)
	fmt.Println("--------------------------------")

	timeKit.ClearInterval(ticker)

	fmt.Println("--------------------------------")
	time.Sleep(time.Second * 11)
	fmt.Println("--------------------------------")
}
