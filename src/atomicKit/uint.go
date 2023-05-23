package atomicKit

import "github.com/gogf/gf/v2/container/gtype"

func NewUint(value ...uint) *gtype.Uint {
	return gtype.NewUint(value...)
}

func NewUint32(value ...uint32) *gtype.Uint32 {
	return gtype.NewUint32(value...)
}

func NewUint64(value ...uint64) *gtype.Uint64 {
	return gtype.NewUint64(value...)
}
