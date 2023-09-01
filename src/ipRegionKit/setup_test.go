package ipRegionKit

import (
	"fmt"
	"github.com/richelieu-yang/chimera/v2/src/consts"
	"github.com/richelieu-yang/chimera/v2/src/core/pathKit"
	"github.com/sirupsen/logrus"
	"testing"
)

func TestGetRegion(t *testing.T) {
	wd, err := pathKit.ReviseWorkingDirInTestMode(consts.ProjectName)
	if err != nil {
		panic(err)
	}
	logrus.Infof("wd: [%s].", wd)

	xdbPath := "chimera-lib/ip2region.xdb"
	MustSetUp(xdbPath)

	//ip := "10.0.9.141"
	ip := "218.90.174.146"
	fmt.Println(GetRegion(ip))
}
