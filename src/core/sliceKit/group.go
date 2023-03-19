package sliceKit

import (
	"github.com/samber/lo"
	lop "github.com/samber/lo/parallel"
)

// Group 分组
/*
PS: 不会修改传参s.

@param s		可以为nil
@param iteratee	不能为nil
@return			保底为空的map实例

e.g.
	s := []int{0, 1, 2, 3, 4, 5}
	m := sliceKit.Group[int, int](s, func(i int) int {
		return i % 3
	})
	fmt.Println(s) // [0 1 2 3 4 5]
	fmt.Println(m) // map[0:[0 3] 1:[1 4] 2:[2 5]]
*/
func Group[T any, U comparable](s []T, iteratee func(item T) U) map[U][]T {
	return lo.GroupBy(s, iteratee)
}

func GroupInParallel[T any, U comparable](s []T, iteratee func(item T) U) map[U][]T {
	return lop.GroupBy(s, iteratee)
}
