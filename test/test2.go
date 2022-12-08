package main

import (
	"github.com/sirupsen/logrus"
	"os"
)

func main() {
	logrus.Infof("[B] value: [%s].", os.Getenv("test"))
}
