package main

import (
	"github.com/gin-gonic/gin"
	"github.com/richelieu-yang/chimera/v2/src/log/logrusKit"
	"github.com/sirupsen/logrus"
)

func main() {
	logrusKit.MustSetUp(&logrusKit.Config{
		Level:      "",
		PrintBasic: false,
	})
	engine := gin.Default()

	engine.Static("/", "./resource1")

	if err := engine.Run(":81"); err != nil {
		logrus.Fatal(err)
	}
}
