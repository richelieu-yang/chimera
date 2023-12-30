package main

import (
	"github.com/richelieu-yang/chimera/v2/src/log/logrusKit"
	"github.com/sirupsen/logrus"
	"golang.org/x/time/rate"
)

func init() {
	logrusKit.MustSetUp(nil)
}

func main() {
	//初始化 limiter 每秒 1 个令牌，令牌桶容量为3
	limiter := rate.NewLimiter(1, 3)
	for i := 0; i < 5; i++ {
		if limiter.Allow() {
			logrus.Infof("%d success", i)
		} else {
			logrus.Infof("%d busy", i)
		}
	}
	logrus.Info("---")

	////阻塞直到获取足够的令牌或者上下文取消
	//ctx, _ := context.WithTimeout(context.Background(), time.Second*2)
	//fmt.Println("start get token", time.Now())
	//err := limiter.WaitN(ctx, 5)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println("success get token", time.Now())
}
