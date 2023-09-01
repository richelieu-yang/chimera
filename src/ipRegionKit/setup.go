package ipRegionKit

import (
	"github.com/lionsoul2014/ip2region/binding/golang/xdb"
	"github.com/richelieu-yang/chimera/v2/src/core/fileKit"
	"github.com/richelieu-yang/chimera/v2/src/log/logrusKit"
	"github.com/sirupsen/logrus"
)

// 缓存整个xdb数据的情况下，searcher对象可以安全用于并发
var searcher *xdb.Searcher

func MustSetUp(xdbPath string) {
	err := SetUp(xdbPath)
	if err != nil {
		logrusKit.DisableQuote(nil)
		logrus.Fatalf("%+v", err)
	}
}

func SetUp(xdbPath string) error {
	var err error
	searcher, err = loadXdbFile(xdbPath)
	return err
}

// GetRegion
/*
PS:
(1) 需要先加载 ip2region.xdb
(2) IP地址查询 - 在线工具（可以查看本机的外网ip）: https://tool.lu/ip/

@return 第1个返回值的格式: 国家|区域|省份|城市|ISP

e.g.
("1.1.1.1")			=> ("澳大利亚|0|0|0|0", nil)
("1.2.3.4") 		=> ("美国|0|华盛顿|0|谷歌", nil)
("10.0.9.141") 		=> ("0|0|0|内网IP|内网IP", nil)
("218.90.174.146") 	=> ("中国|0|江苏省|无锡市|电信", nil)
*/
func GetRegion(ip string) (string, error) {
	// 每个 ip 数据段的 region 信息都固定了格式：国家|区域|省份|城市|ISP，只有中国的数据绝大部分精确到了城市，其他国家部分数据只能定位到国家，后前的选项全部是0。
	return searcher.SearchByStr(ip)
}

// loadXdbFile
/*
@param path xdb文件的路径
*/
func loadXdbFile(xdbPath string) (*xdb.Searcher, error) {
	if err := fileKit.AssertExistAndIsFile(xdbPath); err != nil {
		return nil, err
	}

	// 缓存整个xdb数据
	cBuff, err := xdb.LoadContentFromFile(xdbPath)
	if err != nil {
		return nil, err
	}
	return xdb.NewWithBuffer(cBuff)
}
