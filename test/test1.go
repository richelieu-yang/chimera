package main

import (
	"github.com/richelieu42/go-scales/src/log/zapKit"
	"go.uber.org/zap"
)

func main() {
	path := "logs/c/rocketmq_client_go.log"
	maxFileSize := 1073741824
	maxFileIndex := 10

	logger, err := zapKit.NewSugaredLogger(path, maxFileSize, maxFileIndex, false, zap.InfoLevel)
	if err != nil {
		panic(err)
	}
	logger.Infof("[TEST] %s %d", "ccc", 999)
}
