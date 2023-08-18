package timeKit

import "time"

// SetTimeout
/*
参考：
golang定时器函数 每隔几分钟执行一个函数
	https://www.cnblogs.com/niuben/p/14368715.html
GO语言提前取消定时器
	https://blog.csdn.net/u012265809/article/details/114939168
*/
func SetTimeout(fun func(), duration time.Duration) *time.Timer {
	if fun == nil {
		return nil
	}

	return time.AfterFunc(duration, fun)
}

func ClearTimeout(timer *time.Timer) {
	if timer == nil {
		return
	}

	timer.Stop()
}
