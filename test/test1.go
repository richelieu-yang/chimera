package main

import (
	"github.com/sirupsen/logrus"
	"os"
)

func main() {
	path := os.Getenv("user.home")
	logrus.Info(path)
	logrus.Info(path == "")
}
