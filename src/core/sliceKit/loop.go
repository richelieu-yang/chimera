package sliceKit

import "github.com/samber/lo"
import lop "github.com/samber/lo/parallel"

// ForEach 遍历
/*
@param s		可以为nil
@param iteratee	不能为nil
*/
func ForEach[T any](s []T, iteratee func(item T, index int)) {
	lo.ForEach(s, iteratee)
}

// ForEachInParallel 遍历
/*
@param s		可以为nil
@param iteratee	不能为nil
*/
func ForEachInParallel[T any](s []T, iteratee func(item T, index int)) {
	lop.ForEach(s, iteratee)
}
