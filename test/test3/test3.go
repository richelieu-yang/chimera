package main

import (
	"fmt"
	"golang.org/x/time/rate"
	"time"
)

func main() {
	// 初始化 limiter 每秒10个令牌，令牌桶容量为20
	limiter := rate.NewLimiter(rate.Every(time.Millisecond*100), 20)

	for i := 0; i < 25; i++ {
		if limiter.Allow() {
			fmt.Println(i, "success")
		} else {
			fmt.Println(i, "busy")
		}
	}
	fmt.Println("---")

	time.Sleep(time.Second)
	for i := 0; i < 3; i++ {
		if limiter.Allow() {
			fmt.Println(i, "success")
		} else {
			fmt.Println(i, "busy")
		}
	}
}
