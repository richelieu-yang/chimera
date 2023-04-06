package runeKit

import "unicode"

// IsChineseRune 是否中文字符（包括繁体）？
func IsChineseRune(r rune) bool {
	return unicode.Is(unicode.Han, r)
}
