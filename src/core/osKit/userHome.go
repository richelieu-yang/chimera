package osKit

import "os"

// GetUserHome 获取当前用户目录
/*
@return 可能为""
*/
func GetUserHome() string {
	return os.Getenv("user.home")
}
