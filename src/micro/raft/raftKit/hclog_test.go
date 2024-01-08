package raftKit

import (
	"fmt"
	"github.com/hashicorp/go-hclog"
	"testing"
	"time"
)

func TestNewHcLogger(t *testing.T) {
	fmt.Println(hclog.LevelFromString(""))

	logger := NewHcLogger(&hclog.LoggerOptions{
		Name:   "prefix",
		Level:  hclog.LevelFromString(""),
		Output: hclog.DefaultOutput,
		TimeFn: time.Now,
	})

	//hclog.New(&hclog.LoggerOptions{
	//	Name:  "raft",
	//	Level: hclog.LevelFromString(""),
	//})

	logger.Trace("trace")
	logger.Debug("debug")
	logger.Info("info")
	logger.Warn("warn")
	logger.Error("error")
}
