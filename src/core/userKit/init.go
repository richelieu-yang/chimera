package userKit

import (
	"github.com/richelieu42/go-scales/src/core/errorKit"
	"os/user"
)

var u *user.User
var userHomeDir string

func init() {
	var err error

	u, err = user.Current()
	if err != nil {
		errorKit.Panic("[SCALES] user.Current() fails, error: %+v", err)
	}

	userHomeDir, err = getUserHomeDir()
	if err != nil {
		errorKit.Panic("[SCALES] userHomeDir is invalid, error: %+v", err)
	}
}
