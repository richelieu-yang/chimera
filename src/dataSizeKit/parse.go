package dataSizeKit

import "github.com/dustin/go-humanize"

// ParseString
/*
@return 第一个返回值的单位: 字节(bytes)

e.g.
("42MB")	=> 42000000 <nil>
("42 MB")	=> 42000000 <nil>
("42mib")	=> 44040192 <nil>
("42 mib")	=> 44040192 <nil>
*/
func ParseString(str string) (uint64, error) {
	return humanize.ParseBytes(str)
}
