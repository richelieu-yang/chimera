package etcdKit

import (
	"github.com/richelieu-yang/chimera/v3/src/core/errorKit"
	"github.com/richelieu-yang/chimera/v3/src/core/interfaceKit"
	"github.com/richelieu-yang/chimera/v3/src/core/sliceKit"
)

type (
	Config struct {
		Endpoints []string `json:"endpoints" yaml:"endpoints"`
		// LogPath etcd客户端的日志输出（默认: 输出到控制台）
		LogPath string `json:"logPath,optional" yaml:"logPath"`
	}
)

func (config *Config) Check() error {
	if err := interfaceKit.AssertNotNil(config, "config"); err != nil {
		return err
	}

	config.Endpoints = sliceKit.Uniq(config.Endpoints)
	config.Endpoints = sliceKit.RemoveEmpty(config.Endpoints, true)
	if sliceKit.IsEmpty(config.Endpoints) {
		return errorKit.Newf("config.Endpoints is empty")
	}

	return nil
}
