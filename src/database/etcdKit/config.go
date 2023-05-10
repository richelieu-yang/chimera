package etcdKit

import (
	"github.com/richelieu42/chimera/v2/src/core/errorKit"
	"github.com/richelieu42/chimera/v2/src/core/sliceKit"
)

type (
	Config struct {
		Endpoints []string `json:"endpoints"`
		// LogPath etcd客户端的日志输出（默认: 输出到控制台）
		LogPath string `json:"logPath,optional"`
	}
)

func (config *Config) Check() error {
	if config == nil {
		return errorKit.Simple("config == nil")
	}

	config.Endpoints = sliceKit.Uniq(config.Endpoints)
	config.Endpoints = sliceKit.RemoveEmpty(config.Endpoints, true)
	if len(config.Endpoints) == 0 {
		return errorKit.Simple("config.Endpoints is empty")
	}

	return nil
}
