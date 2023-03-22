package userKit

import (
	"github.com/mitchellh/go-homedir"
)

func GetName() string {
	return u.Name
}

func GetUserName() string {
	return u.Username
}

// GetUserHomeDir 获取当前用户的目录.
/*
@return 必定不为"" && 是个存在的目录

e.g.
() => "/Users/richelieu"
*/
func GetUserHomeDir() string {
	return userHomeDir
}

func getUserHomeDir() (string, error) {
	//// os.Getenv("user.home")可能会返回""，比如在Mac环境下
	//userHomeDir := os.Getenv("user.home")
	//if userHomeDir == "" {
	//	userHomeDir = os.Getenv("HOME")
	//}

	return homedir.Dir()
}
