package userKit

import (
	"github.com/sirupsen/logrus"
	"os/user"
)

var u *user.User
var userHomeDir string

func init() {
	var err error

	u, err = user.Current()
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"error": err.Error(),
		}).Fatal("fail to get current user information")
	}

	userHomeDir = u.HomeDir
	//userHomeDir, err = getUserHomeDir()
	//if err != nil {
	//	logrus.WithFields(logrus.Fields{
	//		"error": err.Error(),
	//	}).Fatal("")
	//
	//	errorKit.Panic("[SCALES] userHomeDir is invalid, error: %+v", err)
	//}
}
