package main

import (
	"github.com/sirupsen/logrus"
	"os/user"
)

func main() {
	u, err := user.Current()
	if err != nil {
		logrus.Fatal(err)
	}
	logrus.Info(u.HomeDir)
}
