package timeKit

import (
	"github.com/richelieu42/chimera/v2/core/sliceKit"
	"time"
)

// ParseStringToTime 类型转换: string => time.Time
/*
@param layout 	时间格式
@param str 		要解析的时间字符串
*/
func ParseStringToTime(layout, str string, args ...*time.Location) (time.Time, error) {
	loc := sliceKit.GetFirstItemWithDefault(nil, args...)
	if loc == nil {
		loc = time.Local
	}

	return time.ParseInLocation(layout, str, loc)
}

// ParseStringToDuration string => time.Duration
/*
@param str e.g. "300ms"、"-1.5h"、"2h45m"
*/
func ParseStringToDuration(str string) (time.Duration, error) {
	return time.ParseDuration(str)
}
