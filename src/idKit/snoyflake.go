package idKit

import (
	"github.com/richelieu42/chimera/v2/src/core/errorKit"
	"github.com/sony/sonyflake"
)

// NewSonyFlake
/*
PS: 可用作分布式唯一id.

golang实现的雪花算法 https://mp.weixin.qq.com/s/visG_GHtU67xCtsvvG1aPQ
*/
func NewSonyFlake(st *sonyflake.Settings) (*sonyflake.Sonyflake, error) {
	if st == nil {
		st = &sonyflake.Settings{}
	}
	sf := sonyflake.NewSonyflake(*st)
	if sf == nil {
		return nil, errorKit.Simple("")
	}
	return sf, nil
}
