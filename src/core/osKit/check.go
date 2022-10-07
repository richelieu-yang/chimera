//go:build !windows

package osKit

import "github.com/richelieu42/go-scales/src/core/errorKit"

const (
	// 最小值
	lowerLimit = 5000
	// 推荐值
	recommendedValue = 65535
)

// CheckEnvironment 检查服务器环境
/*
@return 返回nil代表：通过检查
*/
func CheckEnvironment() error {
	maxOpenFiles, err := GetMaxOpenFiles()
	if err != nil {
		return err
	}
	if maxOpenFiles < lowerLimit {
		return errorKit.Simple("maxOpenFiles(%d) is too small, recommended value is [%d]", maxOpenFiles, recommendedValue)
	}

	userMaxProcesses, err := GetUserMaxProcesses()
	if err != nil {
		return err
	}
	if userMaxProcesses < lowerLimit {
		return errorKit.Simple("userMaxProcesses(%d) is too small, recommended value is [%d]", userMaxProcesses, recommendedValue)
	}

	return nil
}
