package atomicKit

import (
	"github.com/gogf/gf/v2/container/gtype"
)

// NewInterface
/*
PS: 对 atomic.Value 进行了封装.
*/
func NewInterface(value ...interface{}) *gtype.Interface {
	return gtype.NewInterface(value...)
}
