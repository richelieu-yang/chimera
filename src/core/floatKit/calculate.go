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
func Div(x float64, y ...float64) float64 {
	if y == nil {
		return x
	}

	tmp := decimal.NewFromFloat(x)
	for _, p := range y {
		tmp = tmp.Div(decimal.NewFromFloat(p))
	}
	value, _ := tmp.Float64()
	return value
}

// DivAndRound 除法（四舍五入），divides and rounds to a given precision
/*
@param places 保留的小数位

e.g.
(4, 1, 2) =>		0.5
(4, 1, 3) => 		0.3333
(4, 0.123456) => 	0.1235
*/
func DivAndRound(places int32, x float64, y ...float64) float64 {
	tmp := decimal.NewFromFloat(x)
	if y != nil {
		for _, p := range y {
			tmp = tmp.Div(decimal.NewFromFloat(p))
		}
	}

	// 相较于 Div，多了这一行
	tmp = tmp.Round(places)

	value, _ := tmp.Float64()
	return value
}
