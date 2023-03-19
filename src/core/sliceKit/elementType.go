package sliceKit

import "github.com/samber/lo"
import lop "github.com/samber/lo/parallel"

// ConvertElementType 转换切片实例的元素类型
/*
@param string	可以为nil
@param iteratee 不能为nil，会导致panic: runtime error: invalid memory address or nil pointer dereference
@return			必定不为nil（保底空的slice实例）
*/
func ConvertElementType[T any, R any](s []T, iteratee func(item T, index int) R) []R {
	return lo.Map(s, iteratee)
}

// ConvertElementTypeInParallel 转换切片实例的元素类型（使用协程更加高效，但慎用!!!）
/*
@param string	可以为nil
@param iteratee 不能为nil，会导致panic: runtime error: invalid memory address or nil pointer dereference
@return			必定不为nil（保底空的slice实例）
*/
func ConvertElementTypeInParallel[T any, R any](s []T, iteratee func(item T, index int) R) []R {
	return lop.Map(s, iteratee)
}
