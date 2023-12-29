package dataSizeKit

import (
	"github.com/duke-git/lancet/v2/formatter"
)

var (
	// ToReadableSiString SI标准(base 1000)
	/*
		   PS:
		   (1) 采用SI标准（国际单位制；https://blog.csdn.net/bioitee/article/details/120797739）.
		   (2) 1KB == 1000
		   (3) 采用此标准: Mac的访达...

			@param precision 指定小数点后的位数，默认为4
	*/
	ToReadableSiString func(size float64, precision ...int) string = formatter.DecimalBytes

	// ToReadableIecString IEC标准(base 1024)
	/*
		   PS:
		   (1) 采用IEC标准（国际电工委员会；https://blog.csdn.net/bioitee/article/details/120797739）.
		   (2) 1KB == 1024
		   (3) 采用此标准: Windows的文件资源管理器、钉钉（Mac版和Windows版都是）、XManager和Tabby的可视化工具...

			@param precision 指定小数点后的位数，默认为4
	*/
	ToReadableIecString func(size float64, precision ...int) string = formatter.BinaryBytes

	// ParseSiString SI标准(base 1000)
	ParseSiString func(size string) (uint64, error) = formatter.ParseDecimalBytes

	// ParseIecString IEC标准(base 1024)
	ParseIecString func(size string) (uint64, error) = formatter.ParseBinaryBytes
)

func A() {

}
