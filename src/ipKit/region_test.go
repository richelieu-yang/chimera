package ipKit

import (
	"fmt"
	"github.com/lionsoul2014/ip2region/binding/golang/xdb"
	"github.com/richelieu-yang/chimera/v2/src/consts"
	"github.com/richelieu-yang/chimera/v2/src/core/fileKit"
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

	/* 加载 ip2region.xdb 文件 */
	xdbPath := "chimera-lib/ip2region.xdb"
	if err := fileKit.AssertExistAndIsFile(xdbPath); err != nil {
		panic(err)
	}
	// 缓存整个xdb数据
	cBuff, err := xdb.LoadContentFromFile(xdbPath)
	if err != nil {
		panic(err)
	}
	tmpSearcher, err := xdb.NewWithBuffer(cBuff)
	if err != nil {
		panic(err)
	}
	searcher = tmpSearcher

	ip := "218.90.174.146"
	fmt.Println(GetRegion(ip))
}
