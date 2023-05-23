package atomicKit

import "github.com/gogf/gf/v2/container/gtype"

// NewBytes
/*
PS: 对 atomic.Value 进行了封装.
*/
func NewBytes(value ...[]byte) *gtype.Bytes {
	return gtype.NewBytes(value...)
}
