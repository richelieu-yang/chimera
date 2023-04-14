package main

import (
	"github.com/richelieu42/chimera/v2/src/log/logrusKit"
)

func main() {
	logger := logrusKit.NewBasicLogger()
	logger.Info("666")

	//ants.NewPoolWithFunc()
	//
	//pool, err := ants.NewPool(
	//	1,
	//	//ants.WithPanicHandler(func(i interface{}) {
	//	//	logrus.Error(i)
	//	//}),
	//)
	//if err != nil {
	//	logrus.Fatal(err)
	//}
	//
	//err = pool.Submit(func() {
	//	panic("000")
	//})
	//logrus.Info(err)
	//
	//err = pool.Submit(func() {
	//	logrus.Info("111")
	//})
	//logrus.Info(err)
	//
	//select {}
}
