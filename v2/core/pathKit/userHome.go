package pathKit

import (
	"github.com/richelieu42/chimera/v2/core/userKit"
)

func GetUserHomePath() string {
	return userKit.GetUserHomeDir()
}
