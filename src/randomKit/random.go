// Package randomKit
/*
有两个标准库("math/rand"、"crypto/rand")，建议使用第一个，更加全面.
*/
package randomKit

import (
	"github.com/richelieu42/go-scales/src/core/floatKit"
	"math/rand"
	"time"
)

var r *rand.Rand

func init() {
	/*
		参考：https://www.cnblogs.com/jukaiit/p/10785433.html
		不加随机种子，每次遍历获取都是重复的一些随机数据；
		加随机种子，可以保证每次随机都是随机的.
	*/
	//rand.Seed(time.Now().UnixNano())

	r = rand.New(rand.NewSource(time.Now().UnixNano()))

}

// Int
/*
@param n > 0（否则会panic）

@return 返回[0, max)范围内的随机数
*/
func Int(max int) int {
	return r.Intn(max)
}

// Float64
/*
@param places 保留的小数位

@return [0.0, 1.0)
*/
func Float64(places int32) float64 {
	f := r.Float64()
	return floatKit.Floor(f, places)
}

func Bool() bool {
	return Int(2) == 1
}
