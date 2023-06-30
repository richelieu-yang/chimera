package runeKit

import "unicode"

var (
	// IsLetter 判断是否为字母
	IsLetter func(r rune) bool = unicode.IsLetter

	// IsDigit 判断是否为数字
	IsDigit func(r rune) bool = unicode.IsDigit

	// IsSpace 判断是否为空白符号('\t', '\n', '\v', '\f', '\r', ' ', U+0085 (NEL), U+00A0 (NBSP))
	IsSpace func(r rune) bool = unicode.IsSpace
)

// IsChineseRune 是否中文字符（包括繁体）？
func IsChineseRune(r rune) bool {
	return unicode.Is(unicode.Han, r)
}
