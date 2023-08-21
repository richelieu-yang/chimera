//go:build go1.21

package mathKit

import "golang.org/x/exp/constraints"

// Max
/*
PS: 使用 Go1.21 新增的内置函数 max.
*/
func Max[T constraints.Ordered](x T, y ...T) T {
	return max[T](x, y...)
}

// Min
/*
PS: 使用 Go1.21 新增的内置函数 min.
*/
func Min[T constraints.Ordered](x T, y ...T) T {
	return min[T](x, y...)
}
