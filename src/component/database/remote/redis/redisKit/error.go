package redisKit

import (
	"github.com/richelieu-yang/chimera/v2/src/core/errorKit"
)

var (
	NotSetupError = errorKit.New("Haven’t been set up correctly")
)
