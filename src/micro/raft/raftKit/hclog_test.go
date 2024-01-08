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
	logger.Info("info")
	logger.Warn("warn")
	logger.Error("error")
}
