// Package strKit
/*
空白包括: " "（空格）、\r、\n、\t、\f
*/
package strKit

import (
	"regexp"
	"strings"
)

// TrimSpace 去掉左右的空格（包括：\r、\n），不包括中间的空格.
/*
PS: Golang中所有的传参都是值传递（传值），都是一个副本，一个拷贝.
*/
var TrimSpace func(str string) string = strings.TrimSpace

// RemoveSpace 移除str中所有的（包括左右的、中间的）：" "、"\t"、"\r"、"\n"...
/*
e.g.
("    \t\r\n  \n\n") => ""
*/
func RemoveSpace(str string) string {
	/*
		\s: 空白 (相当于 [\t\n\f\r ])
		x+: 匹配一个或多个 x，优先匹配更多(贪婪)
	*/
	reg := regexp.MustCompile("\\s+")
	return reg.ReplaceAllString(str, "")
}

// ReplaceSpacesWithSpace 将连续的空格替换为单个空格
/*
e.g.
"It\tis\na\fsimple\rtest                 !" => "It is a simple test !"
*/
func ReplaceSpacesWithSpace(str string) string {
	reg := regexp.MustCompile("\\s+")
	return reg.ReplaceAllString(str, " ")
}
