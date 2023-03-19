package mapKit

import "github.com/samber/lo"

// MapToSlice map实例 => slice实例
/*
@param m
@param iteratee 传参: 一对键值，返回值: slice实例中的一个元素
*/
func MapToSlice[K comparable, V any, R any](m map[K]V, iteratee func(key K, value V) R) []R {
	return lo.MapToSlice(m, iteratee)
}
