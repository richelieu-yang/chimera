package main

import (
	"github.com/hashicorp/go-hclog"
	"time"
)

func main() {
	logger := hclog.New(&hclog.LoggerOptions{
		Name:   "test",
		Level:  hclog.DefaultLevel,
		Output: hclog.DefaultOutput,
		TimeFn: time.Now,
	})
	logger.Info("123")

	//hclog.Default().SetLevel(hclog.Debug)
	//
	//var logger hclog.Logger
	//
	//logger := &Logger{}

}
