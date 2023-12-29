package dataSizeKit

import "github.com/dustin/go-humanize"

var (
	// ParseString
	/*
	   @param str 同时支持"IEC标准"和"SI标准"
	   @return 第一个返回值的单位: 字节(bytes)

	   e.g.
	   ("42MB")	=> 42000000 <nil>
	   ("42 MB")	=> 42000000 <nil>
	   ("42mib")	=> 44040192 <nil>
	   ("42 mib")	=> 44040192 <nil>
	*/
	ParseString func(str string) (uint64, error) = humanize.ParseBytes

	// ToReadableStringWithIEC
	/*
	   PS:
	   (1) 采用IEC标准（国际电工委员会；https://blog.csdn.net/bioitee/article/details/120797739）.
	   (2) 1KB == 1024
	   (3) 采用此标准: Windows的文件资源管理器、钉钉（Mac版和Windows版都是）、XManager和Tabby的可视化工具...

	   @param s 单位: 字节(bytes)

	   e.g.
	   	(78848) => "77 KiB"
	*/
	ToReadableStringWithIEC func(s uint64) string = humanize.IBytes

	// ToReadableStringWithSI
	/*
	   PS:
	   (1) 采用SI标准（国际单位制；https://blog.csdn.net/bioitee/article/details/120797739）.
	   (2) 1KB == 1000
	   (3) 采用此标准: Mac的访达...

	   @param s 单位: 字节(bytes)

	   e.g.
	   	(78848) => "79 kB"
	*/
	ToReadableStringWithSI func(s uint64) string = humanize.Bytes
)
