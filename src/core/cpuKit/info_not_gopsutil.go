//go:build !(darwin || windows || (linux && 386) || (linux && amd64) || (linux && arm) || (linux && arm64))

package cpuKit

import (
	"github.com/sirupsen/logrus"
)

func printBasicDetails(logger *logrus.Logger) {

}
