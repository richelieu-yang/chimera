package sliceKit

import "github.com/samber/lo"

// Compact 去除零值.
/*
@param s 可以为nil
@return 必定不为nil（保底为空的slice实例）

e.g.

*/
func Compact[T comparable](s []T) []T {
	return lo.Compact(s)
}
