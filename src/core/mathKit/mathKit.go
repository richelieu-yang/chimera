package mathKit

import "math"

// Round 四舍五入
/*
TODO: 看后续版本Golang math.Round() 的返回值是否与别的语言一致.

！！！：
(1) 此方法与"Java（Math.round）"和"JavaScript（Math.round）"一致；
(2) Golang 1.17.7原生的math.Round()与"Java（Math.round()）"和"JavaScript（Math.round()）"不一致，传参为负数的情况.
*/
func Round(x float64) float64 {
	// 思路：+0.5；向下取整.
	return math.Floor(x + 0.5)
}

// Ceil 向上取整
/*
PS: 个人感觉，向右取整数.

e.g.
(3.6) 	=>	4
(-3.6) 	=>	-3
*/
func Ceil(x float64) float64 {
	return math.Ceil(x)
}

// Floor 向下取整
/*
PS: 个人感觉，向左取整数.

e.g.
(3.6) 	=>	3
(-3.6) 	=>	-4
*/
func Floor(x float64) float64 {
	return math.Floor(x)
}
