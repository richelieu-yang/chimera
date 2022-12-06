package userKit

import (
	"github.com/mitchellh/go-homedir"
	"github.com/richelieu42/go-scales/src/core/file/fileKit"
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
	userHomeDir, err := homedir.Dir()
	if err != nil {
		return "", err
	}
	if err := fileKit.AssertExistAndIsDir(userHomeDir); err != nil {
		return "", err
	}
	return userHomeDir, nil
}
