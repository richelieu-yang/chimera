package mathKit

import (
	"github.com/samber/lo"
	"golang.org/x/exp/constraints"
)

// Sum 求和
/*
e.g.
([]int{0, 1, 2, 3}) => 6
*/
func Sum[T constraints.Float | constraints.Integer | constraints.Complex](s []T) T {
	return lo.Sum(s)
}

// SumBy 求和
/*
@param s 		可以为nil（此时返回0）
@param iteratee	(1) 不能为nil（除非s == nil），否则会导致panic: runtime error: invalid memory address or nil pointer dereference
				(2) 传参为T类型，返回值为R类型
e.g.
	i := mathKit.SumBy[string, int]([]string{"0", "1", "2"}, func(item string) int {
		tmp, _ := strconv.Atoi(item)
		return tmp
	})
	fmt.Println(i) // 3
*/
func SumBy[T any, R constraints.Float | constraints.Integer | constraints.Complex](s []T, iteratee func(item T) R) R {
	return lo.SumBy(s, iteratee)
}
