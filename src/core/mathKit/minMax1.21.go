// go版本 >= 1.21
//go:build go1.21

package mathKit

import "cmp"

func Max[T cmp.Ordered](x T, y ...T) T {
	return max(x, y...)
}

func Min[T cmp.Ordered](x T, y ...T) T {
	return min(x, y...)
}
