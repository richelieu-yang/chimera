package timeKit

import (
	"time"
)

// Parse 类型转换: string => time.Time
/*
PS:
(1) 为什么不直接使用 time.Parse()？
	因为time.Parse使用 time.UTC 作为loc，会有时差.
(2) 本函数使用 time.Local 作为loc.

@param layout 	时间格式
@param str 		要解析的时间字符串

e.g.
(timeKit.FormatDate, "2016-08-08")
*/
func Parse[F ~string](format F, timeStr string) (time.Time, error) {
	loc := time.Local
	return time.ParseInLocation(string(format), timeStr, loc)
}

// ParseInLocation
/*
@param loc time.Local || time.UTC
*/
func ParseInLocation[F ~string](format F, timeStr string, loc *time.Location) (time.Time, error) {
	return time.ParseInLocation(string(format), timeStr, loc)
}

// ParseDuration string => time.Duration
/*
@param str (1) 如果为 ""，将返回error
    	   (2) e.g. "300ms"、"-1.5h"、"2h45m"
*/
var ParseDuration func(str string) (time.Duration, error) = time.ParseDuration
