package mathKit

import "github.com/duke-git/lancet/v2/mathutil"

func MinBy[T any](slice []T, comparator func(T, T) bool) T {
	return mathutil.MinBy(slice, comparator)
}

func MaxBy[T any](slice []T, comparator func(T, T) bool) T {
	return mathutil.MaxBy(slice, comparator)
}
