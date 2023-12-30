package main

import (
	"fmt"
	"golang.org/x/time/rate"
	"time"
)

//func RateLimitMiddleware(rate int) gin.HandlerFunc {
//	rl := ratelimit.New(rate) // 每秒rate个请求
//
//	return func(c *gin.Context) {
//		_ = rl.Take()
//		c.Next()
//	}
//}

func main() {
	//初始化 limiter 每秒生成1个令牌，令牌桶容量为20
	limiter := rate.NewLimiter(rate.Every(time.Second), 1)
	//模拟单位时间执行多次操作
	for i := 0; i < 5; i++ {
		if limiter.Allow() {
			fmt.Println("发送邮件")
		} else {
			fmt.Println("请求多次，过滤")
		}
	}
	if limiter.Allow() {
		fmt.Println("发送邮件")
	}
}
