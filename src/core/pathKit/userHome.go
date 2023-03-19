package pathKit

import (
	"github.com/richelieu42/chimera/src/core/userKit"
)

func GetUserHomePath() string {
	return userKit.GetUserHomeDir()
}
