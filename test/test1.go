package main

import (
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func main() {
	engine := gin.Default()
	pprof.Register(engine)
	if err := engine.Run(":8080"); err != nil {
		logrus.Fatal(err)
	}
}
