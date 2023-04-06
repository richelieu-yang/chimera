package timeKit

import (
	"time"
)

// SetInterval
/*
参考：
golang定时器函数 每隔几分钟执行一个函数	https://www.cnblogs.com/niuben/p/14368715.html
*/
func SetInterval(fun func(t time.Time), duration time.Duration) *time.Ticker {
	if fun == nil {
		return nil
	}

	ticker := time.NewTicker(duration)
	go func() {
		for tt := range ticker.C {
			fun(tt)
		}
	}()
	return ticker
}

func ClearInterval(ticker *time.Ticker) {
	if ticker == nil {
		return
	}
	ticker.Stop()
}
