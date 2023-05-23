package atomicKit

import "github.com/gogf/gf/v2/container/gtype"

func NewInt(value ...int) *gtype.Int {
	return gtype.NewInt(value...)
}

func NewInt32(value ...int32) *gtype.Int32 {
	return gtype.NewInt32(value...)
}

func NewInt64(value ...int64) *gtype.Int64 {
	return gtype.NewInt64(value...)
}
