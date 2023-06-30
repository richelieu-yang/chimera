package strKit

import "strings"

var (
	// Trim 返回字符串，删除了s"左边和右边"的连续cutset
	/*
		e.g.
		fmt.Println(strings.Trim("aaa0aaa0aaa", "a")) // "0aaa0"
		fmt.Println(strings.Trim(" aba ", "a"))       // " aba "
		fmt.Println(strings.Trim("/a/", "/"))         // "a"
	*/
	Trim func(s, cutset string) string = strings.Trim

	// TrimLeft 返回字符串，删除了s"左边"的连续cutset
	/*
		e.g.
		fmt.Println(strings.TrimLeft("aaa0aaa0aaa", "a")) // "0aaa0aaa"
		fmt.Println(strings.TrimLeft(" aba ", "a"))       // " aba "
	*/
	TrimLeft func(s, cutset string) string = strings.TrimLeft

	// TrimRight 返回字符串，删除了s"右边"的连续cutset
	/*
		e.g.
		fmt.Println(strings.TrimRight("aaa0aaa0aaa", "a")) // "aaa0aaa0"
		fmt.Println(strings.TrimRight(" aba ", "a"))       // " aba "
	*/
	TrimRight func(s, cutset string) string = strings.TrimRight

	// TrimFunc
	TrimFunc func(s string, f func(rune) bool) string = strings.TrimFunc

	// TrimLeftFunc
	TrimLeftFunc func(s string, f func(rune) bool) string = strings.TrimLeftFunc

	// TrimRightFunc
	TrimRightFunc func(s string, f func(rune) bool) string = strings.TrimRightFunc
)
