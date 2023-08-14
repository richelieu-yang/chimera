package timeKit

import (
	"time"
)

// Parse 类型转换: string => time.Time
/*
@param layout 	时间格式
@param str 		要解析的时间字符串

e.g.
(timeKit.FormatDate, "2016-08-08")
*/
func Parse[T ~string](format T, timeStr string) (time.Time, error) {
	return time.Parse(string(format), timeStr)
}

func ParseInLocation[T ~string](format T, timeStr string, loc *time.Location) (time.Time, error) {
	return time.ParseInLocation(string(format), timeStr, loc)
}

// ParseDuration string => time.Duration
/*-
@param str (1) 如果为 ""，将返回error
    	   (2) e.g. "300ms"、"-1.5h"、"2h45m"
*/
var ParseDuration func(str string) (time.Duration, error) = time.ParseDuration
