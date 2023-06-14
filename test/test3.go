package main

import (
	"github.com/sirupsen/logrus"
	"time"
)

func main() {
	var t time.Duration = time.Second * 119
	logrus.Infof("%s", t)
}
