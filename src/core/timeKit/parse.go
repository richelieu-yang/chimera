package timeKit

import (
	"github.com/richelieu42/chimera/v2/src/core/sliceKit"
	"time"
)

// ParseTimeString 类型转换: string => time.Time
/*
@param layout 	时间格式
@param str 		要解析的时间字符串
*/
func ParseTimeString(layout, str string, args ...*time.Location) (time.Time, error) {
	loc := sliceKit.GetFirstItemWithDefault(nil, args...)
	if loc == nil {
		loc = time.Local
	}

	return time.ParseInLocation(layout, str, loc)
}

// ParseDurationString string => time.Duration
/*
@param str e.g."300ms"、"-1.5h"、"2h45m"
*/
func ParseDurationString(str string) (time.Duration, error) {
	return time.ParseDuration(str)
}
