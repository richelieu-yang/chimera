//go:build !(darwin || windows || (linux && 386) || (linux && amd64) || (linux && arm) || (linux && arm64))

package diskKit

import (
	"github.com/richelieu-yang/chimera/v2/src/core/errorKit"
	"github.com/sirupsen/logrus"
)

func GetDiskUsageStatsByPath(path string) (*DiskUsageStats, error) {
	return nil, errorKit.New("Currently not supported")
}

func PrintBasicDetails(logger *logrus.Logger) {

}
