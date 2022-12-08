package main

import (
	"github.com/sirupsen/logrus"
	"os"
)

func main() {
	if err := os.Setenv("test", "测试"); err != nil {
		panic(err)
	}
	logrus.Infof("value: [%s].", os.Getenv("test"))
}
