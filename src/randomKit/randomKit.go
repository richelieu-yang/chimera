// Package randomKit
/*
有两个标准库("math/rand"、"crypto/rand")，建议使用第一个，更加全面.
*/
package randomKit

import (
	"github.com/richelieu42/go-scales/src/core/errorKit"
	"github.com/richelieu42/go-scales/src/core/floatKit"
	"math/rand"
	"time"
)

func init() {
	/*
		参考：https://www.cnblogs.com/jukaiit/p/10785433.html
		不加随机种子，每次遍历获取都是重复的一些随机数据；
		加随机种子，可以保证每次随机都是随机的.
	*/
	rand.Seed(time.Now().UnixNano())
}

func Int(max int) int {
	i, _ := IntE(max)
	return i
}

// IntE
/*
@return 返回[0, max)范围内的随机数
*/
func IntE(max int) (int, error) {
	if max <= 0 {
		return 0, errorKit.Simple("max(%d) is invalid", max)
	}
	rst := rand.Intn(max)
	return rst, nil
}

func Float64(max float64, places int32) float64 {
	f, _ := Float64E(max, places)
	return f
}

// Float64E
/*
@param places 小数位
@return 返回[0, max)范围内的随机数（有可能返回0）
*/
func Float64E(max float64, places int32) (float64, error) {
	if max <= 0 {
		return 0.0, errorKit.Simple("max(%d) is invalid", max)
	}

	tmp := rand.Float64()
	if max != 1.0 {
		tmp *= max
	}
	return floatKit.Round(tmp, places), nil
}

// Bool
/*
TODO: 有性能问题，待优化
*/
func Bool() bool {
	return rand.Int31n(2) == 1
}
