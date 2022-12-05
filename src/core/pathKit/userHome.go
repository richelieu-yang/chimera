package pathKit

import (
	"github.com/richelieu42/go-scales/src/core/osKit"
)

func GetUserHomePath() string {
	return osKit.GetUserHome()
}
