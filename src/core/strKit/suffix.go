package strKit

import "strings"

// EndWith
/*
PS: 区分大小写.

@param s		suffix不为""的情况下，如果s为""，返回值必定为false
@param suffix 	如果suffix为""，返回值必定为true

e.g.	""的情况
("", "") 	=> true
("1", "")	=> true
("", "1")	=> false

e.g.1	区分大小写
("abc", "abc") => true
("abc", "Abc") => false
*/
var EndWith func(s, suffix string) bool = strings.HasSuffix

var CutSuffix func(s, suffix string) (before string, found bool) = strings.CutSuffix

// RemoveSuffixIfExists 去掉指定的"后缀"（如果存在的话）
/*
PS:
(1) 区分大小写；
(2) 存在多个的话，只会移除最后1个.
*/
var RemoveSuffixIfExists func(s, suffix string) string = strings.TrimSuffix

// AppendIfMissing 如果给定字符串不是以给定的字符串为结尾，则在"尾部"添加结尾字符串（不忽略大小写）.
/*
PS: 区分大小写.

@param suffix 后缀

e.g.
("abc", "c"))	=> "abc"
("abc", "C")) 	=> "abcC"
("abc", "0")) 	=> "abc0"
*/
func AppendIfMissing(str, suffix string) (rst string) {
	if EndWith(str, suffix) {
		rst = str
	} else {
		rst = str + suffix
	}
	return
}
