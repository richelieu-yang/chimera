package pathKit

import (
	"github.com/richelieu42/go-scales/src/core/userKit"
)

func GetUserHomePath() string {
	return userKit.GetUserHomeDir()
}
