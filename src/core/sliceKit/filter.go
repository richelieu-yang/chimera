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

// FilterAndRevise 过滤（可额外处理满足条件的元素）.
/*
@param s		可以为nil
@param callback (1) 不能为nil
				(2) 第2个返回值: 是否满足过滤条件？
				(3) 第2个返回值 == true的情况下，将第1个返回值加到返回的slice实例中.
@return 必定不为nil（保底为空的slice实例）

e.g.
	s := sliceKit.FilterAndRevise([]string{"cpu", "gpu", "mouse", "keyboard"}, func(item string, index int) (string, bool) {
		if strings.HasSuffix(item, "pu") {
			return "right-" + item, true
		}
		return "", false
	})
	fmt.Println(s) // [right-cpu right-gpu]
*/
func FilterAndRevise[T any, R any](s []T, callback func(item T, index int) (R, bool)) []R {
	return lo.FilterMap(s, callback)
}
