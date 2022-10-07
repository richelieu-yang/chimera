package sliceKit

import (
	"github.com/richelieu42/go-scales/src/core/strKit"
)

// ToString 切片实例（元素任意类型） => string
/*
参考: https://mp.weixin.qq.com/s/tvy9L-pb_8WFWAmA9u-bMg

e.g.
([]int{-1, 0, 1}) => "[-1 0 1]"
*/
func ToString[T any](s []T) string {
	return strKit.Format("%v", s)
}
