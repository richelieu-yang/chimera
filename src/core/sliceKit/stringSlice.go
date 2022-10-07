package sliceKit

import (
	"gitee.com/richelieu042/go-scales/src/core/strKit"
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

// TrimAndRemoveEmpty 对每个元素：先trim，为空就从切片中移除掉
/*
@param s 可以为 nil 或长度为0的切片实例
@return nil || 新的[]string实例（不会修改传参s）

e.g.
(nil) 						=> nil
([]string{""}) 				=> []string{}
([]string{"1", "", " 2 "}) 	=> []string{"1", "2"}
*/
func TrimAndRemoveEmpty(s []string) []string {
	if s == nil {
		return nil
	}

	rst := make([]string, 0, len(s))
	for _, str := range s {
		str = strKit.Trim(str)
		if strKit.IsNotEmpty(str) {
			rst = append(rst, str)
		}
	}
	return rst
}
