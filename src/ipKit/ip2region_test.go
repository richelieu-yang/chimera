package ipKit

import (
	"gitee.com/richelieu042/go-scales/src/core/pathKit"
	"github.com/sirupsen/logrus"
	"testing"
)

func TestGetRegionByIp(t *testing.T) {
	if err := pathKit.ReviseProjectDirWhenTesting(); err != nil {
		panic(err)
	}

	region, err := GetRegionByIp("218.90.174.146")
	if err != nil {
		panic(err)
	}
	logrus.Infof("region: [%s].", region)
}
