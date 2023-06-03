package etcdKit

import (
	"github.com/richelieu-yang/chimera/v2/src/core/errorKit"
	"github.com/richelieu-yang/chimera/v2/src/core/sliceKit"
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
		return errorKit.New("config == nil")
	}

	config.Endpoints = sliceKit.Uniq(config.Endpoints)
	config.Endpoints = sliceKit.RemoveEmpty(config.Endpoints, true)
	if len(config.Endpoints) == 0 {
		return errorKit.New("config.Endpoints is empty")
	}

	return nil
}
