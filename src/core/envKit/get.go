package envKit

import (
	"os"
)

var (
	GetEnv func(key string) string = os.Getenv
)

// GetEnvWithDefault （带默认值地）获取环境变量.
func GetEnvWithDefault(key string, def string) (rst string) {
	rst = GetEnv(key)

	if rst == "" {
		rst = def
	}
	return
}
