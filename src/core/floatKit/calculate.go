// Package floatKit
/*
浮点数的高精度运算.
*/
package floatKit

import (
	"github.com/shopspring/decimal"
)

// Add 加
func Add(delta float64, ps ...float64) float64 {
	if ps == nil {
		return delta
	}

	tmp := decimal.NewFromFloat(delta)
	for _, p := range ps {
		tmp = tmp.Add(decimal.NewFromFloat(p))
	}
	value, _ := tmp.Float64()
	return value
}

// Sub 减
func Sub(delta float64, ps ...float64) float64 {
	if ps == nil {
		return delta
	}

	tmp := decimal.NewFromFloat(delta)
	for _, p := range ps {
		tmp = tmp.Sub(decimal.NewFromFloat(p))
	}
	value, _ := tmp.Float64()
	return value
}

// Mul 乘
func Mul(delta float64, ps ...float64) float64 {
	if ps == nil {
		return delta
	}

	tmp := decimal.NewFromFloat(delta)
	for _, p := range ps {
		tmp = tmp.Mul(decimal.NewFromFloat(p))
	}
	value, _ := tmp.Float64()
	return value
}

// Div 除法（不额外处理小数位）
func Div(delta float64, ps ...float64) float64 {
	if ps == nil {
		return delta
	}

	tmp := decimal.NewFromFloat(delta)
	for _, p := range ps {
		tmp = tmp.Div(decimal.NewFromFloat(p))
	}
	value, _ := tmp.Float64()
	return value
}

// DivRound 除法（四舍五入），divides and rounds to a given precision
/*
@param precision 小数位数

e.g.
(4, 1, 2) =>		0.5
(4, 1, 3) => 		0.3333
(4, 0.123456) => 	0.1235
*/
func DivRound(precision int32, delta float64, ps ...float64) float64 {
	tmp := decimal.NewFromFloat(delta)
	if ps != nil {
		for _, p := range ps {
			tmp = tmp.Div(decimal.NewFromFloat(p))
		}
	}
	tmp = tmp.Round(precision)
	value, _ := tmp.Float64()
	return value
}

// Round 保留小数位（四舍五入），类似于 math.Round()，但功能更强大
/*
PS:
(1) 参考：https://zhuanlan.zhihu.com/p/152050239?from_voters_page=true
(2) 个人感觉：先把正负号拿出来 => 进行取舍 => 把正负号还回去
(3) 传参为负数的情况下，Golang的四舍五入与别的语言（Java、JavaScript）不同，详情见"Golang.docx"中的"math标准库".

@param places 小数位数（如果最后几个都是0的话，会省略掉）；可以为负值

e.g.
(3.14, 1)	=> 3.1
(3.15, 1)	=> 3.2
(-3.14, 1) 	=> -3.1
(-3.15, 1) 	=> -3.2

(3.1001, 2) => 3.1
(521, -1) 	=> 520
*/
func Round(f float64, places int32) float64 {
	f, _ = decimal.NewFromFloat(f).Round(places).Float64()
	return f
}

// Ceil 向上取整，类似于 math.Ceil()，但功能更强大
/*
PS:
个人感觉: x轴向右.

e.g.
(3.14, 1)	=> 3.2
(-3.14, 1)	=> -3.1
*/
func Ceil(f float64, places int32) float64 {
	f, _ = decimal.NewFromFloat(f).RoundCeil(places).Float64()
	return f
}

// Floor 向下取整，类似于 math.Floor()，但功能更强大
/*
PS:
个人感觉: x轴向左.

e.g.
(3.14, 1)	=> 3.1
(-3.14, 1)	=> -3.2
*/
func Floor(f float64, places int32) float64 {
	f, _ = decimal.NewFromFloat(f).RoundFloor(places).Float64()
	return f
}
