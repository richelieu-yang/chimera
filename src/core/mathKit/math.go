package mathKit

import (
	"github.com/duke-git/lancet/v2/mathutil"
	"github.com/samber/lo"
	"golang.org/x/exp/constraints"
)

// Clamp clamps number within the inclusive lower and upper bounds.
/*
case value < min: 	返回min
case value > max: 	返回max
case others:		返回value

e.g.
(0, -10, 10)	=> 0
(-42, -10, 10)	=> -10
(42, -10, 10)	=> 10
*/
func Clamp[T constraints.Ordered](value T, min T, max T) T {
	return lo.Clamp(value, min, max)
}

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

// Exponent 指数.
/*
@return x^n

e.g.
	rst := mathKit.Exponent(2, 10)
	fmt.Println(rst) // 1024
*/
var Exponent func(x, n int64) int64 = mathutil.Exponent
