package mapKit

import (
	"github.com/richelieu42/go-scales/src/core/boolKit"
	"github.com/richelieu42/go-scales/src/core/intKit"
	"github.com/richelieu42/go-scales/src/core/strKit"
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

// GetBool
/*
@param m 可以为nil
*/
func GetBool[K comparable, V any](m map[K]V, key K) (bool, error) {
	value := Get(m, key)
	return boolKit.ToBoolE(value)
}
