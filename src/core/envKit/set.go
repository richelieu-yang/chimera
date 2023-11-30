package envKit

import "os"

var (
	SetEnv func(key, value string) error = os.Setenv
)

// SetEnvs 批量设置环境变量.
func SetEnvs(m map[string]string) (err error) {
	for k, v := range m {
		if err = os.Setenv(k, v); err != nil {
			return
		}
	}
	return
}
