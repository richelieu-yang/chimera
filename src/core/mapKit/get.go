package mapKit

import (
	"github.com/richelieu-yang/chimera/v2/src/core/boolKit"
	"github.com/richelieu-yang/chimera/v2/src/core/intKit"
	"github.com/richelieu-yang/chimera/v2/src/core/strKit"
)

// Get
/*
@param m 	可以为nil
@param key	如果不存在对应的值，将返回值类型的零值
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

// GetInt32
/*
@param m 可以为nil
*/
func GetInt32[K comparable, V any](m map[K]V, key K) (int32, error) {
	value := Get(m, key)
	return intKit.ToInt32E(value)
}

// GetInt64
/*
@param m 可以为nil
*/
func GetInt64[K comparable, V any](m map[K]V, key K) (int64, error) {
	value := Get(m, key)
	return intKit.ToInt64E(value)
}

// GetBool
/*
@param m 可以为nil
*/
func GetBool[K comparable, V any](m map[K]V, key K) (bool, error) {
	value := Get(m, key)
	return boolKit.ToBoolE(value)
}
