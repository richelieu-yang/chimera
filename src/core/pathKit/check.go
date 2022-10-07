package pathKit

import "gitee.com/richelieu042/go-scales/src/core/strKit"

// CheckSkip 检查是否发生"路径穿越"（路径穿透）
/*
@param path	 	父路径
@param path1	子路径
@return true: 发生"路径穿越"

e.g.
("/a//b/", "/a//b//../c.docx")	=> true
("/a//b/", "//a//b///c.docx")	=> false
*/
func CheckSkip(parent, path string) bool {
	parent = Join(parent)
	path = Join(path)

	return !strKit.StartWith(path, parent)
}
