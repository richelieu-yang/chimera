package componentKit

import (
	"github.com/richelieu42/go-scales/src/core/pathKit"
	"testing"
)

func Test(t *testing.T) {
	if err := pathKit.ReviseProjectDirWhenTesting(); err != nil {
		panic(err)
	}

	if err := InitializeEnvironment(); err != nil {
		panic(err)
	}
}
