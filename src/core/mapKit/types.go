package mapKit

import "github.com/samber/lo"

// MapKeys 修改map实例的 键类型（key）
func MapKeys[K comparable, V any, R comparable](in map[K]V, iteratee func(value V, key K) R) map[R]V {
	return lo.MapKeys(in, iteratee)
}

// MapValues 修改map实例的 值类型（value）
/*
PS:
(1) 不会修改传参in;
(2) 返回的是一个新的map实例.

@return 必定不为nil
*/
func MapValues[K comparable, V any, R any](in map[K]V, iteratee func(value V, key K) R) map[K]R {
	return lo.MapValues(in, iteratee)
}

// MapEntries 修改map实例的 键类型（key）和值类型（value）
func MapEntries[K1 comparable, V1 any, K2 comparable, V2 any](in map[K1]V1, iteratee func(key K1, value V1) (K2, V2)) map[K2]V2 {
	return lo.MapEntries(in, iteratee)
}
