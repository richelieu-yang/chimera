package skyWalkingKit

import (
	"github.com/richelieu-yang/chimera/v3/src/core/errorKit"
	"github.com/richelieu-yang/chimera/v3/src/netKit"
	"time"
)

type (
	Config struct {
		ServerAddr    string        `json:"serverAddr"`
		CheckInterval time.Duration `json:"checkInterval"`
	}
)

func (config *Config) Verify() error {
	if config == nil {
		return errorKit.New("config == nil")
	}

	addr, err := netKit.ParseToAddress(config.ServerAddr)
	if err != nil {
		return err
	}
	config.ServerAddr = addr.String()

	if config.CheckInterval <= 0 {
		config.CheckInterval = time.Second
	}

	return nil
}
