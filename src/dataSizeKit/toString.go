package dataSizeKit

import "github.com/dustin/go-humanize"

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
func ToReadableStringWithIEC(s uint64) string {
	return humanize.IBytes(s)
}

// ToReadableStringWithSI
/*
Deprecated: 大部门场景使用IEC标准: ToReadableStringWithIEC().

PS:
(1) 采用SI标准（国际单位制；https://blog.csdn.net/bioitee/article/details/120797739）.
(2) 1KB == 1000
(3) 采用此标准: Mac的访达...

@param s 单位: 字节(bytes)

e.g.
	(78848) => "79 kB"
*/
func ToReadableStringWithSI(s uint64) string {
	return humanize.Bytes(s)
}
