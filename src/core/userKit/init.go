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
		logrus.Fatal(err)
	}
	userHomeDir = u.HomeDir
}
