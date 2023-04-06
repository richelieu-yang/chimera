package strKit

import (
	"unicode"
	"unicode/utf8"
)

// GetRuneCount
/*
PS:
(1) 包括简体、繁体.
(2) 如果确定 传参str 中不存在中文字符，建议直接使用 内置函数len().

@param str （可能）带有 中文字符 的字符串

e.g.
("test")	=> 4
("测试")		=> 2
*/
func GetRuneCount(str string) int {
	//return len([]rune(str))
	return utf8.RuneCountInString(str)
}

// GetChineseRuneCount 获取字符串中 中文字符 的个数
/*
PS: 包括简体、繁体.
*/
func GetChineseRuneCount(str string) (count int) {
	// 此处 r 的类型为 int32（rune）
	for _, r := range str {
		if unicode.Is(unicode.Han, r) {
			count++
		}
	}
	return
}

// HasChineseRune 字符串中是否存在 中文字符？
/*
PS: 包括简体、繁体.
*/
func HasChineseRune(str string) bool {
	// 此处 r 的类型为 int32（rune）
	for _, r := range str {
		if unicode.Is(unicode.Han, r) {
			return true
		}
	}
	return false
}
