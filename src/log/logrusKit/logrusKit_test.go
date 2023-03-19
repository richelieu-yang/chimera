package logrusKit

import (
	"github.com/richelieu42/chimera/src/core/pathKit"
	"testing"
)

func TestInitialize(t *testing.T) {
	if err := pathKit.ReviseProjectDirWhenTesting(); err != nil {
		panic(err)
	}

	InitializeByDefault()
	PrintBasicDetails()
}
