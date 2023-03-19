package sliceKit

import "github.com/samber/lo"
import lop "github.com/samber/lo/parallel"

// ConvertElementType 转换切片实例的元素类型
/*
@param string	可以为nil
@param iteratee 不能为nil，会导致panic: runtime error: invalid memory address or nil pointer dereference
@return			必定不为nil（保底空的slice实例）

e.g. []int => []string
	s := sliceKit.ConvertElementType([]int{0, 1, 2, 3}, func(item int, index int) string {
		return "0x" + strconv.Itoa(item)
	})
	fmt.Println(s) // [0x0 0x1 0x2 0x3]
*/
func ConvertElementType[T any, R any](s []T, iteratee func(item T, index int) R) []R {
	return lo.Map(s, iteratee)
}

// ConvertElementTypeInParallel 转换切片实例的元素类型
/*
PS: 使用协程更加高效，但要慎用!!!

@param string	可以为nil
@param iteratee 不能为nil，会导致panic: runtime error: invalid memory address or nil pointer dereference
@return			必定不为nil（保底空的slice实例）
*/
func ConvertElementTypeInParallel[T any, R any](s []T, iteratee func(item T, index int) R) []R {
	return lop.Map(s, iteratee)
}
