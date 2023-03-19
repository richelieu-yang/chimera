package sliceKit

import (
	"github.com/samber/lo"
	lop "github.com/samber/lo/parallel"
)

// Group 分组
/*
@param s		可以为nil
@param iteratee	不能为nil
@return			保底为空的map实例
*/
func Group[T any, U comparable](s []T, iteratee func(item T) U) map[U][]T {
	return lo.GroupBy(s, iteratee)
}

func GroupInParallel[T any, U comparable](s []T, iteratee func(item T) U) map[U][]T {
	return lop.GroupBy(s, iteratee)
}
