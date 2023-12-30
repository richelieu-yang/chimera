package rateLimitKit

import "golang.org/x/time/rate"

var (
	// NewLimiter
	/*
		令牌桶算法.

		@param r 每秒生成的令牌数，(1) 可以直接是 某个数值;
								(2) 也可以是 rate.Every() 的返回值;
		@param b 令牌桶的容量

		e.g.
		// 初始化 limiter 每秒10个令牌，令牌桶容量为20
		limiter := rate.NewLimiter(rate.Every(time.Millisecond*100), 20)
	*/
	NewLimiter func(r rate.Limit, b int) *rate.Limiter = rate.NewLimiter
)
