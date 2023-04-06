package sliceKit

import "github.com/samber/lo"

// Count
/*
e.g.
	count := lo.Count([]int{1, 5, 1}, 1)
	// 2
*/
func Count[T comparable](collection []T, value T) (count int) {
	return lo.Count(collection, value)
}

// CountBy
/*
e.g.
	count := lo.CountBy([]int{1, 5, 1}, func(i int) bool {
		return i < 4
	})
	// 2
*/
func CountBy[T any](collection []T, predicate func(item T) bool) (count int) {
	return lo.CountBy(collection, predicate)
}

// CountValues 计算集合中每个元素的个数.
/*
e.g.
	lo.CountValues([]int{1, 2, 2})
	// map[int]int{1: 1, 2: 2}
*/
func CountValues[T comparable](collection []T) map[T]int {
	return lo.CountValues(collection)
}

// CountValuesBy
/*
e.g.
	isEven := func(v int) bool {
		return v%2==0
	}
	lo.CountValuesBy([]int{1, 2, 2}, isEven)
	// map[bool]int{false: 1, true: 2}
*/
func CountValuesBy[T any, U comparable](collection []T, mapper func(item T) U) map[U]int {
	return lo.CountValuesBy(collection, mapper)
}
