package mapKit

// ContainsKey 判断 map实例 中是否存在 指定的key
/*
@param m 可以为nil（此时返回值固定为false）

e.g.
(m, "1") => false
*/
func ContainsKey[K comparable, V any](m map[K]V, key K) bool {
	_, ok := m[key]
	return ok
}

// ContainsKeys 判断 map实例 中是否存在 所有指定的key
/*
@param keys 可以一个key都不传，此时将固定返回true
*/
func ContainsKeys[K comparable, V any](m map[K]V, keys ...K) bool {
	for _, key := range keys {
		if !ContainsKey(m, key) {
			return false
		}
	}
	return true
}

// GetKeySlice 获取map实例中的所有key
/*
@param m 	如果为 nil 或 空的map实例，将返回nil
@return 	非nil的slice实例（len >= 0）
*/
func GetKeySlice[K comparable, V any](m map[K]V) []K {
	s := make([]K, 0, len(m))

	for key := range m {
		s = append(s, key)
	}
	return s
}
