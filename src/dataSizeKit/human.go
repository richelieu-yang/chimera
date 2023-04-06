package dataSizeKit

import "github.com/dustin/go-humanize"

// ToReadableStringWithSI
/*
PS:
(1) 采用SI标准（国际单位制；https://blog.csdn.net/bioitee/article/details/120797739）.
(2) 1KB == 1000；

@param s 单位: 字节(bytes)

e.g. 该文件在Mac的文件系统中的大小是79 KB
	(78848) => "79 kB"
*/
func ToReadableStringWithSI(s uint64) string {
	return humanize.Bytes(s)
}

// ToReadableStringWithIEC
/*
PS:
(1) 采用IEC标准（国际电工委员会；https://blog.csdn.net/bioitee/article/details/120797739）.
(2) 1KB == 1024

@param s 单位: 字节(bytes)
*/
func ToReadableStringWithIEC(s uint64) string {
	return humanize.IBytes(s)
}
