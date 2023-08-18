package logrusKit

import "time"

// ToZeroAM
/*
@return 传参t对应的当天凌晨零点
*/
func ToZeroAM(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
}
