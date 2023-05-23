package atomicKit

import "github.com/gogf/gf/v2/container/gtype"

// NewString
/*
PS: 对 atomic.Value 进行了封装.
*/
func NewString(value ...string) *gtype.String {
	return gtype.NewString(value...)
}
