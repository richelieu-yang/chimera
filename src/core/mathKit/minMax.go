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

// Max
/*
Deprecated: Go1.21使用内置函数 max().
*/
func Max[T constraints.Ordered](p T, args ...T) (max T) {
	max = p

	for _, ele := range args {
		if ele > max {
			max = ele
		}
	}
	return max
}

// Min
/*
Deprecated: Go1.21使用内置函数 min().
*/
func Min[T constraints.Ordered](p T, args ...T) (min T) {
	rst := p

	for _, ele := range args {
		if ele < rst {
			rst = ele
		}
	}
	return rst
}
