package idKit

import (
	"github.com/richelieu-yang/chimera/v2/src/core/errorKit"
	"github.com/sony/sonyflake"
)

// NewSonyFlake 雪花算法.
/*
PS:
(1) 可用作分布式唯一id（前提是合理配置好 MachineID）;
(2) 通过 sonyflake.Sonyflake.NextID() 生成id（貌似是18位的）.

golang实现的雪花算法 https://mp.weixin.qq.com/s/visG_GHtU67xCtsvvG1aPQ
*/
func NewSonyFlake(st *sonyflake.Settings) (*sonyflake.Sonyflake, error) {
	if st == nil {
		st = &sonyflake.Settings{}
	}
	sf := sonyflake.NewSonyflake(*st)
	if sf == nil {
		return nil, errorKit.New("sf == nil")
	}
	return sf, nil
}
