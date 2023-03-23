package redisKit

import (
	"github.com/richelieu42/chimera/src/assertKit"
	"github.com/richelieu42/chimera/src/core/errorKit"
)

func MustSetUp(config *Config) {
	err := SetUp(config)
	assertKit.Must(err)
}

func SetUp(config *Config) error {
	if config == nil {
		return errorKit.Simple("config == nil")
	}

	return nil
}
