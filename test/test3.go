package main

import (
	"github.com/sirupsen/logrus"
)

func main() {
	// 程序正常退出不会触发 logrus.RegisterExitHandle() 注册的exit handlers，因此加上此defer语句
	defer logrus.Fatal("Exit normally.")

	logrus.RegisterExitHandler(func() {
		logrus.Info(111)
	})
}
