package timeKit

import "time"

// ParseStringToTime 类型转换: string => time.Time
/*
@param layout 	时间格式
@param str 		时间字符串
*/
func ParseStringToTime(layout, str string) (time.Time, error) {
	return time.Parse(layout, str)
}

func ParseStringToTimeWithLocation(layout, str string, loc *time.Location) (time.Time, error) {
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
