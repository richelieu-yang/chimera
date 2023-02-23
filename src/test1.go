package main

import (
	"github.com/richelieu42/go-scales/src/log/logrusKit"
	"github.com/sirupsen/logrus"
)

func main() {
	logger, err := logrusKit.NewFileLogger("aaa.txt", nil, logrus.DebugLevel, false)
	if err != nil {
		panic(err)
	}

	logger.Info(666)

	//for i := 0; i < 3; i++ {
	//	logger.Info(i)
	//	time.Sleep(time.Second)
	//}
	//
	//go func() {
	//	logger.Panic("ym")
	//}()
	//
	//select {}
}
