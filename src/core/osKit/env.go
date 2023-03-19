package osKit

import (
	"github.com/richelieu42/chimera/src/core/sliceKit"
	"os"
)

// GetEnv （带默认值地）获取系统环境变量.
/*
@param 	可以为""
@return 可能为""

e.g.
("") 			=> ""
("JAVA_HOME") 	=> "/Library/Java/JavaVirtualMachines/zulu-8.jdk/Contents/Home"
*/
func GetEnv(key string, defArgs ...string) string {
	val := os.Getenv(key)

	if val == "" {
		// 默认值
		def := sliceKit.GetFirstItemWithDefault("", defArgs...)
		return def
	}
	return val
}

// SetEnvs 一次性设置多个系统变量.
/*
@param m 可以为nil
@return 如果发生error的话，返回第一个
*/
func SetEnvs(m map[string]string) error {
	for k, v := range m {
		if err := os.Setenv(k, v); err != nil {
			return err
		}
	}
	return nil
}
