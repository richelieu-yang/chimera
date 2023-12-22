package main

import "github.com/sirupsen/logrus"

func main() {
	l := logrus.GetLevel()
	logrus.Info(l)
}
