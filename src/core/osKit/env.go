package osKit

import (
	"github.com/richelieu42/go-scales/src/core/sliceKit"
	"os"
)

// GetEnv （带默认值地）获取系统环境变量
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
		return sliceKit.GetFirstItemWithDefault("", defArgs...)
	}
	return val
}
