package sliceKit

import "github.com/samber/lo"

// Replace
/*
@param collection 	可以为nil
@param n			替换的个数
					(1) > 0，替换指定个数
					(2) = 0，一个都不替换
					(3) < 0，替换所有
@return 保底为length==0的切片实例
*/
func Replace[T comparable](collection []T, old T, new T, n int) []T {
	return lo.Replace(collection, old, new, n)
}

// ReplaceAll 替换所有.
/*
@param collection 	可以为nil
@return 保底为length==0的切片实例
*/
func ReplaceAll[T comparable](collection []T, old T, new T) []T {
	return lo.ReplaceAll(collection, old, new)
}
