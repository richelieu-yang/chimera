package etcdKit

import "github.com/richelieu42/chimera/v2/src/core/errorKit"

var (
	NotSetupError = errorKit.New("hasn't been set up")
)
