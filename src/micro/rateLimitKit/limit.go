package rateLimitKit

import "golang.org/x/time/rate"

var (
	// NewLimiter 限流器，基于令牌桶算法（Token Bucket）.
	/*
		PS: 初始状态下，令牌桶是满的.

		@param r 每秒生成的令牌数，(1) 可以直接是 某个数值; (2) 也可以是 rate.Every() 的返回值;
		@param b 令牌桶的容量

		e.g.
		// 初始化 limiter 每秒10个令牌，令牌桶容量为20
		limiter := rate.NewLimiter(rate.Every(time.Millisecond*100), 20)
	*/
	NewLimiter func(limit rate.Limit, burst int) *rate.Limiter = rate.NewLimiter
)
