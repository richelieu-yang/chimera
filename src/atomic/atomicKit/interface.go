package atomicKit

import (
	"github.com/gogf/gf/v2/container/gtype"
)

// NewInterface
/*
PS: 对 atomic.Value 进行了封装.
*/
var NewInterface func(value ...interface{}) *gtype.Interface = gtype.NewInterface
