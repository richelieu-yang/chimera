package sliceKit

import "github.com/samber/lo"

// Contains
/*
@param s 可以为nil（返回false）
*/
func Contains[T comparable](s []T, element T) bool {
	return lo.Contains(s, element)
}

// ContainsBy
/*
@param s			可以为nil（返回false）
@param predicate 	不能为nil
*/
func ContainsBy[T any](s []T, predicate func(item T) bool) bool {
	return lo.ContainsBy(s, predicate)
}
