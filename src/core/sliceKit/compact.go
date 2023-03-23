package sliceKit

import "github.com/samber/lo"

// Compact 去除零值.
/*
@param s 可以为nil；不会修改传参s
@return 必定不为nil（保底为空的slice实例）

e.g.
	s := []string{"", "foo", "", "bar", ""}
	s1 := sliceKit.Compact[string](s)
	fmt.Println(s1) 	// [foo bar]
*/
func Compact[T comparable](s []T) []T {
	return lo.Compact(s)
}
