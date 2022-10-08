package componentKit

import (
	"github.com/richelieu42/go-scales/src/core/pathKit"
	"github.com/richelieu42/go-scales/src/core/strKit"
	"github.com/sirupsen/logrus"
	"testing"
)

func Test(t *testing.T) {
	if err := pathKit.ReviseProjectDirWhenTesting(); err != nil {
		panic(err)
	}

	if err := InitializeEnvironment(); err != nil {
		panic(err)
	}

	if err := InitializeRedisComponent(); err != nil {
		panic(err)
	}

	msgProcessor := func(code string, msg string, data interface{}) string {
		return strKit.Format("[%s] %s", code, msg)
	}
	if err := InitializeJsonComponent(msgProcessor, ""); err != nil {
		panic(err)
	}

	logrus.Info("----------------------------------")
}
