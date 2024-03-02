package userKit

import (
	"github.com/sirupsen/logrus"
	"os/user"
)

var u *user.User

func init() {
	var err error

	u, err = user.Current()
	if err != nil {
		logrus.Fatal("Fail to get current user, error: ", err.Error())
	}
}

// GetUid user ID
func GetUid() string {
	return u.Uid
}

// GetGid primary group ID
func GetGid() string {
	return u.Gid
}

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
	return u.HomeDir
}

//func getUserHomeDir() (string, error) {
//	//// os.Getenv("user.home")可能会返回""，比如在Mac环境下
//	//userHomeDir := os.Getenv("user.home")
//	//if userHomeDir == "" {
//	//	userHomeDir = os.Getenv("HOME")
//	//}
//
//	return homedir.Dir()
//}
