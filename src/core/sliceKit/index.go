package sliceKit

import "github.com/samber/lo"

// IndexOf
/*
@param s 	可以为nil（此时返回-1）
@return 	如果不存在于切片实例中的话，返回-1
*/
func IndexOf[T comparable](s []T, element T) int {
	return lo.IndexOf(s, element)
}

// LastIndexOf
/*
@param s 	可以为nil（此时返回-1）
@return 	如果不存在于切片实例中的话，返回-1
*/
func LastIndexOf[T comparable](s []T, element T) int {
	return lo.LastIndexOf(s, element)
}
