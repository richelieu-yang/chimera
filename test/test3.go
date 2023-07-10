package main

import (
	"github.com/sirupsen/logrus"
)

func main() {
	i := 0
	for {
		i++
		logrus.Info(i)
	}
}
