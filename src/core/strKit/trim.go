package strKit

import "strings"

var (
	// Trim 返回字符串，删除了s"左边&&右边"的连续cutset
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

	// TrimFunc 返回字符串，删除了s"左边&&右边"的连续的满足条件的rune
	/*
		@param f	(1) 返回true: 	删除此rune，继续判断下一个rune
					(2) 返回false: 	不删除此rune，中断此次（左边||右边）删除
	*/
	TrimFunc func(s string, f func(rune) bool) string = strings.TrimFunc

	// TrimLeftFunc 返回字符串，删除了s"左边"的连续的满足条件的rune
	/*
		@param f	(1) 返回true: 	删除此rune，继续判断下一个rune
					(2) 返回false: 	不删除此rune，中断此次（左边||右边）删除
	*/
	TrimLeftFunc func(s string, f func(rune) bool) string = strings.TrimLeftFunc

	// TrimRightFunc 返回字符串，删除了s"右边"的连续的满足条件的rune
	/*
		@param f	(1) 返回true: 	删除此rune，继续判断下一个rune
					(2) 返回false: 	不删除此rune，中断此次（左边||右边）删除
	*/
	TrimRightFunc func(s string, f func(rune) bool) string = strings.TrimRightFunc
)
