package raftKit

import (
	"github.com/hashicorp/go-hclog"
	"testing"
	"time"
)

func TestNewHcLogger(t *testing.T) {
	logger := NewHcLogger(&hclog.LoggerOptions{
		Name:   "prefix",
		Level:  hclog.DefaultLevel,
		Output: hclog.DefaultOutput,
		TimeFn: time.Now,
	})

	logger.Debug("debug")
	logger.Info("info")   // 2023-12-27T14:45:57.214+0800 [INFO]  prefix: info
	logger.Warn("warn")   // 2023-12-27T14:45:57.215+0800 [WARN]  prefix: warn
	logger.Error("error") // 2023-12-27T14:45:57.215+0800 [ERROR] prefix: error
}
