package mapKit

import (
	"github.com/richelieu42/go-scales/src/core/boolKit"
	"github.com/richelieu42/go-scales/src/core/intKit"
	"github.com/richelieu42/go-scales/src/core/strKit"
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

// ContainKeys 判断 map实例 中是否存在 所有指定的key
/*
@param m 可以为nil（此时返回值固定为false）
@param keys 可以一个key都不传，此时将固定返回true
*/
func ContainKeys[K comparable, V any](m map[K]V, keys ...K) bool {
	for _, key := range keys {
		if !Contains(m, key) {
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

// Get
/*

@param m 	可以为nil
@param key	如果不存在与 传参key 对应的值，将返回值类型的零值
*/
func Get[K comparable, V any](m map[K]V, key K) V {
	return m[key]
}

// GetString
/*
@param m 可以为nil
*/
func GetString[K comparable, V any](m map[K]V, key K) (string, error) {
	value := Get(m, key)
	return strKit.ToStringE(value)
}

// GetInt
/*
@param m 可以为nil
*/
func GetInt[K comparable, V any](m map[K]V, key K) (int, error) {
	value := Get(m, key)
	return intKit.ToIntE(value)
}

// GetBool
/*
@param m 可以为nil
*/
func GetBool[K comparable, V any](m map[K]V, key K) (bool, error) {
	value := Get(m, key)
	return boolKit.ToBoolE(value)
}
