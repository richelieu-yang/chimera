package runtimeKit

import (
	"github.com/richelieu42/chimera/v2/core/errorKit"
	"github.com/richelieu42/chimera/v2/dataSizeKit"
	"runtime/debug"
)

// SetSoftMemoryLimit 软内存限制(soft memory limit)
/*
PS:
(1) This limit will be respected even if GOGC=off (or, if SetGCPercent(-1) is executed).
(2) 默认memory limit是math.MaxInt64（即第一次调用debug.SetMemoryLimit()的返回值为math.MaxInt64）.

@param limit 单位: B(byte；字节)
@return 之前的内存限制
*/
func SetSoftMemoryLimit(limit int64) (int64, error) {
	if limit <= 0 {
		return 0, errorKit.Simple("limit(%d) is invalid", limit)
	}
	if limit < int64(512*dataSizeKit.MB.GetValue()) {
		return 0, errorKit.Simple("limit(%s) is too small", dataSizeKit.ToReadableDataSizeString(uint64(limit)))
	}
	return debug.SetMemoryLimit(limit), nil
}
