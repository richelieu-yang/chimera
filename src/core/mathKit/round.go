package mathKit

import "github.com/duke-git/lancet/v2/mathutil"

// Round 四舍五入，保留n位小数.
/*
@param n 保留的小数位（可以 < 0，依然生效）

e.g.
	(123.124, -1)	// 120
	(123.124, 0)	// 123
	(123.124, 1)	// 123.1
	(100.125, 2)	// 100.13
	(100.125, 3)	// 100.125
*/
var Round func(x float64, n int) float64 = mathutil.RoundToFloat

// RoundToString 四舍五入，保留n位小数.
/*
@param n 保留的小数位（可以 < 0，依然生效）
*/
var RoundToString func(x float64, n int) string = mathutil.RoundToString

// TruncRound 截断n位小数（不进行四舍五入）
/*
@param n 保留的小数位（可以 < 0，但有点奇怪!!!）

e.g.
(1234.124, 0)	=> 1234
(1234.124, -1)	=> 1234
(1234.124, -2)	=> 0

(100.125, 2)	=> 100.12
(100.125, 3)	=> 100.125
*/
var TruncRound func(x float64, n int) float64 = mathutil.TruncRound
