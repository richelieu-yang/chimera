package main

import (
	"context"
	"github.com/richelieu-yang/chimera/v2/src/log/logrusKit"
	"github.com/sirupsen/logrus"
	"golang.org/x/time/rate"
	"time"
)

func init() {
	logrusKit.MustSetUp(nil)
}

func main() {
	// 初始化 limiter: 每秒 1 个令牌，令牌桶容量为 3
	limiter := rate.NewLimiter(1, 3)

	// (1) 先清空令牌桶
	for i := 0; i < 5; i++ {
		if limiter.Allow() {
			logrus.Infof("%d success", i)
		} else {
			logrus.Infof("%d busy", i)
		}
	}
	logrus.Info("---")

	// (2) 阻塞直到获取足够的令牌 或者 上下文取消
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	logrus.Info("get token starts")
	err := limiter.WaitN(ctx, 3)
	if err != nil {
		panic(err)
	}
	logrus.Info("get token ends")
}
