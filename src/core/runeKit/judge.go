package runeKit

import "unicode"

var (
	// IsLetter 判断是否为（大写||小写）字母？
	/*
		e.g.
			fmt.Println(runeKit.IsLetter('A')) // true
			fmt.Println(runeKit.IsLetter('a')) // true
	*/
	IsLetter func(r rune) bool = unicode.IsLetter

	// IsDigit 判断是否为数字？
	/*
		e.g.
			fmt.Println(runeKit.IsDigit('1')) // true
			fmt.Println(runeKit.IsDigit('0')) // true
			fmt.Println(runeKit.IsDigit('-')) // false
	*/
	IsDigit func(r rune) bool = unicode.IsDigit

	// IsSpace 判断是否为空白符号('\t', '\n', '\v', '\f', '\r', ' ', U+0085 (NEL), U+00A0 (NBSP))？
	/*
		e.g.
			fmt.Println(runeKit.IsSpace(' '))  // true
			fmt.Println(runeKit.IsSpace('\r')) // true
			fmt.Println(runeKit.IsSpace('\n')) // true
			fmt.Println(runeKit.IsSpace('\t')) // true
	*/
	IsSpace func(r rune) bool = unicode.IsSpace
)

// IsChineseRune 是否为（简体||繁体）中文字符？
/*
e.g.
	fmt.Println(runeKit.IsChineseRune('体')) // true
	fmt.Println(runeKit.IsChineseRune('體')) // true
*/
func IsChineseRune(r rune) bool {
	return unicode.Is(unicode.Han, r)
}
