package mathKit

import (
	"github.com/duke-git/lancet/v2/mathutil"
	"github.com/samber/lo"
	"golang.org/x/exp/constraints"
)

// Average 计算平均数.
/*
PS: 可能需要对返回值进行四舍五入.
*/
func Average[T constraints.Integer | constraints.Float](numbers ...T) T {
	return mathutil.Average[T](numbers...)
}

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

// Exponent 指数.
/*
@return x^n

e.g.
	rst := mathKit.Exponent(2, 10)
	fmt.Println(rst) // 1024
*/
var Exponent func(x, n int64) int64 = mathutil.Exponent

// Abs 绝对值.
func Abs[T constraints.Integer | constraints.Float](x T) T {
	return mathutil.Abs(x)
}
