package sliceKit

import "github.com/samber/lo"

// Filter 过滤.
/*
@param s			(1) 可以为nil
					(2) 不会修改传参s
@param predicate	(1) true: 该元素加到返回的slice实例中
					(2) 不能为nil，会导致panic: runtime error: invalid memory address or nil pointer dereference
@return 必定不为nil（保底为空的slice实例）

e.g.
	s := sliceKit.Filter([]int{0, 1, 2, 3}, func(item int, index int) bool {
		return item >= 2
	})
	fmt.Println(s) // [2 3]
*/
func Filter[V any](s []V, predicate func(item V, index int) bool) []V {
	return lo.Filter(s, predicate)
}

// FilterMap
/*
@param s		可以为nil
@param callback 不能为nil
@return 必定不为nil（保底为空的slice实例）
*/
func FilterMap[T any, R any](s []T, callback func(item T, index int) (R, bool)) []R {
	return lo.FilterMap(s, callback)
}
