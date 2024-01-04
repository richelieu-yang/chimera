//go:build !(386 || amd64 || arm || arm64)

package cpuKit

import (
	"github.com/sirupsen/logrus"
)

func printBasicDetails(logger *logrus.Logger) {

}
