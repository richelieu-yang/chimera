package mathKit

import (
	"github.com/duke-git/lancet/v2/mathutil"
	"golang.org/x/exp/constraints"
)

func MinBy[T any](slice []T, comparator func(T, T) bool) T {
	return mathutil.MinBy(slice, comparator)
}

func MaxBy[T any](slice []T, comparator func(T, T) bool) T {
	return mathutil.MaxBy(slice, comparator)
}

// Max 获取最大值.
/*
PS: Go1.21即以上，建议使用内置函数 max().
*/
func Max[T constraints.Ordered](p T, args ...T) (max T) {
	max = p

	for _, ele := range args {
		if ele > max {
			max = ele
		}
	}
	return
}

// Min 获取最小值.
/*
PS: Go1.21即以上，建议使用内置函数 min().
*/
func Min[T constraints.Ordered](p T, args ...T) (min T) {
	min = p

	for _, ele := range args {
		if ele < min {
			min = ele
		}
	}
	return
}
