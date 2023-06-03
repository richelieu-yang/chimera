package pathKit

import (
	"github.com/richelieu-yang/chimera/v2/src/core/userKit"
)

func GetUserHomePath() string {
	return userKit.GetUserHomeDir()
}
