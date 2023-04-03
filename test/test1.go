package main

import (
	"github.com/richelieu42/chimera/src/confKit"
	"github.com/richelieu42/chimera/src/ginKit"
	"github.com/sirupsen/logrus"
)

type (
	Config struct {
		Gin *ginKit.Config
	}
)

func main() {
	config := &Config{}
	confKit.MustLoad("chimera-lib/env.yaml", config)

	logrus.Info(config)
}
