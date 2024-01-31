package gtypeKit

import "github.com/gogf/gf/v2/container/gtype"

// NewString
/*
PS: 对 atomic.Value 进行了封装.
*/
var NewString func(value ...string) *gtype.String = gtype.NewString
