package mathKit

import (
	"github.com/duke-git/lancet/v2/mathutil"
	"github.com/shopspring/decimal"
)

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

//// Round 四舍五入，保留n位小数.
///*
//@param n 保留的小数位（可以 < 0，依然生效）
//
//e.g.
//	(123.124, -1)	// 120
//	(123.124, 0)	// 123
//	(123.124, 1)	// 123.1
//	(100.125, 2)	// 100.13
//	(100.125, 3)	// 100.125
//*/
//var Round func(x float64, n int) float64 = mathutil.RoundToFloat

//// RoundToString 四舍五入，保留n位小数.
///*
//PS: 相较于 Round ，多了最后一步转换为string类型.
//*/
//var RoundToString func(x float64, n int) string = mathutil.RoundToString

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

// Ceil 向上取整，类似于 math.Ceil()，但功能更强大.
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

// Floor 向下取整，类似于 math.Floor()，但功能更强大.
/*
PS:
(1) 个人感觉: x轴向左.

e.g.
(3.14, 1)	=> 3.1
(-3.14, 1)	=> -3.2
*/
func Floor(f float64, places int32) float64 {
	f, _ = decimal.NewFromFloat(f).RoundFloor(places).Float64()
	return f
}
