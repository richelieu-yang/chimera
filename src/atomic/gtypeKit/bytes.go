package gtypeKit

import "github.com/gogf/gf/v2/container/gtype"

// NewBytes
/*
PS: 对 atomic.Value 进行了封装.
*/
var NewBytes func(value ...[]byte) *gtype.Bytes = gtype.NewBytes
