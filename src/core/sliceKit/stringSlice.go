package sliceKit

import (
	"github.com/richelieu42/chimera/src/core/strKit"
	"strings"
)

// Join []string => string
/*
@param sep 分隔符

e.g.
(nil, "-") 							=> ""
([]string{}, "-") 					=> ""
([]string{"0", "1", "2", "3"}, "-") => "0-1-2-3"
*/
func Join(s []string, sep string) string {
	return strings.Join(s, sep)
}

// RemoveEmpty
/*
@param trimArgs 是否 先 对每个元素进行trim操作？默认：false

e.g.
(nil) 			=> nil
([]string{""})	=> []string{}
*/
func RemoveEmpty(s []string, trimArgs ...bool) []string {
	if s == nil {
		return nil
	}

	trimFlag := GetFirstItemWithDefault(false, trimArgs...)

	rst := make([]string, 0, len(s))
	if trimFlag {
		for _, str := range s {
			str = strKit.Trim(str)
			if strKit.IsNotEmpty(str) {
				rst = append(rst, str)
			}
		}
	} else {
		for _, str := range s {
			//str = strKit.Trim(str)
			if strKit.IsNotEmpty(str) {
				rst = append(rst, str)
			}
		}
	}
	return rst
}

// ContainsStringIgnoreCase 字符串str是否在切片s中？（不区分大小写）
func ContainsStringIgnoreCase(s []string, str string) bool {
	for _, tmp := range s {
		if strKit.EqualsIgnoreCase(tmp, str) {
			return true
		}
	}
	return false
}
