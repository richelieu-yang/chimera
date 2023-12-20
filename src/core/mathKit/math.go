package mathKit

import (
	"github.com/duke-git/lancet/v2/mathutil"
	"github.com/samber/lo"
	"golang.org/x/exp/constraints"
)

var (
	// Exponent 指数计算（x的n次方）
	/*
	   @return x^n

	   e.g.
	   	(2, 10)	=> 1024
	   	(8, 2)	=> 64
	*/
	Exponent func(x, n int64) int64 = mathutil.Exponent

	// Factorial 计算阶乘
	Factorial func(x uint) uint = mathutil.Factorial

	// Fibonacci 计算斐波那契数列的第n个数
	Fibonacci func(first, second, n int) int = mathutil.Fibonacci
)

// Average 计算平均数（可能需要对结果调用RoundToFloat方法四舍五入）
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

// Abs 绝对值.
func Abs[T constraints.Integer | constraints.Float](x T) T {
	return mathutil.Abs(x)
}

// Sin 正弦函数.
/*
Deprecated: 有bug（传参precision无效），不建议使用.
*/
var Sin func(radian float64, precision ...int) float64 = mathutil.Sin

// Cos 余弦函数.
var Cos func(radian float64, precision ...int) float64 = mathutil.Cos
