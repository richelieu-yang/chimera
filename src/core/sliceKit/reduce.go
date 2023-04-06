package sliceKit

import "github.com/samber/lo"

// Reduce 将slice实例归纳为一个值（从左到右遍历）.
/*
PS: 排除骚操作，不会修改传参s.

@param s 			可以为nil
@param accumulator	不能为nil
@param initial		初始值
*/
func Reduce[T any, R any](s []T, accumulator func(agg R, item T, index int) R, initial R) R {
	return lo.Reduce(s, accumulator, initial)
}

// ReduceRight 将slice实例归纳为一个值（从右到左遍历；和 Reduce 相反）.
/*
PS: 排除骚操作，不会修改传参s.

@param s 			可以为nil
@param accumulator	不能为nil
@param initial		初始值
*/
func ReduceRight[T any, R any](s []T, accumulator func(agg R, item T, index int) R, initial R) R {
	return lo.ReduceRight(s, accumulator, initial)
}
