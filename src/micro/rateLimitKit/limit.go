package rateLimitKit

import "golang.org/x/time/rate"

var (
	// NewLimiter
	/*
		令牌桶算法.
	*/
	NewLimiter func(r rate.Limit, b int) *rate.Limiter = rate.NewLimiter
)
