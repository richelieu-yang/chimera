package mapKit

import (
	"github.com/samber/lo"
)

// Contains 判断 map实例 中是否存在 指定的key.
/*
@param m 可以为nil（此时返回值固定为false）

e.g.
	(map[string]interface{}(nil), "1") => false
*/
func Contains[K comparable, V any](m map[K]V, key K) bool {
	_, ok := m[key]
	return ok
}

// ContainKeys 判断 map实例 中是否存在 所有指定的key.
/*
@param m 	可以为nil（此时返回值固定为 false）
@param keys (1) 可以一个key都不传，此时将固定返回 true;
			(2) 多个key的情况下，只要有1个key不存在于 传参m 中，将返回false.
*/
func ContainKeys[K comparable, V any](m map[K]V, keys ...K) bool {
	for _, key := range keys {
		if !Contains(m, key) {
			return false
		}
	}
	return true
}

// Remove
/*
PS:
(1) 可能会修改传参m（移除条目的话），因为它是map类型；
(2) 命名参考了 java.util.Map .

@param m 	可以为nil
@param key	可以在传参m中不存在
@return 被移除出map的条目的值（存在的话） + 传参m是否包含传参key

e.g.
	m := map[string]interface{}{
		"a": 0,
		"b": 1,
		"c": 2,
	}

	fmt.Println(mapKit.Remove(m, "b")) // 1 true
	fmt.Println(m)                     // map[a:0 c:2]
*/
func Remove[K comparable, V any](m map[K]V, key K) (V, bool) {
	value, exist := m[key]
	if exist {
		// 存在的话，移除对应条目
		delete(m, key)
	}
	return value, exist
}

// Set 设置值（或更新值）
/*
@param m 不能为nil（否则会导致 panic: assignment to entry in nil map）
*/
func Set[K comparable, V any](m map[K]V, key K, value V) {
	m[key] = value
}

// SetSafely 设置值（或更新值）
/*
@param m 可以为nil
@return 可能是一个新的map实例
*/
func SetSafely[K comparable, V any](m map[K]V, key K, value V) map[K]V {
	if m == nil {
		m = make(map[K]V)
	}

	m[key] = value
	return m
}

// SetSafelyAndHandleOldValue 设置值（或更新值） && 处理旧的值（有的话）
func SetSafelyAndHandleOldValue[K comparable, V any](m map[K]V, key K, value V, handler func(v V) error) error {
	var oldValue V
	var oldExist bool

	if m == nil {
		m = make(map[K]V)
	} else {
		oldValue, oldExist = m[key]
	}

	m[key] = value
	if oldExist {
		return handler(oldValue)
	}
	return nil
}

// Keys creates an array of the map keys.
/*
@param m	可以为nil
@return 	保底空的slice实例
*/
func Keys[K comparable, V any](m map[K]V) []K {
	return lo.Keys(m)
}

// Values creates an array of the map values.
/*
@param m	可以为nil
@return 	保底空的slice实例
*/
func Values[K comparable, V any](m map[K]V) []V {
	return lo.Values(m)
}

// Invert 倒置（反转，交换key、value，交换键和值）
/*
PS: 如果map包含重复值，则后续值将覆盖前一个值的属性赋值。

@param in 可以为nil（将返回空的map实例）
@return 必定不为nil

e.g.
	m1 := lo.Invert(map[string]int{"a": 1, "b": 2})
	fmt.Println(m1) // map[1:a 2:b]
	m2 := lo.Invert(map[string]int{"a": 1, "b": 2, "c": 1})
	fmt.Println(m2) // map[1:c 2:b] 或 map[1:a 2:b]（因为map是无序的）
*/
func Invert[K comparable, V comparable](in map[K]V) map[V]K {
	return lo.Invert(in)
}
