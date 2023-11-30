package timeKit

import "time"

// ToZeroAM
/*
@return 传参t对应的当天凌晨零点

e.g.
	now := time.Now()
	fmt.Println(now) 	// 2023-08-18 15:24:03.167655 +0800 CST m=+0.004041126
	t := timeKit.ToZeroAM(now)
	fmt.Println(t) 		// 2023-08-18 00:00:00 +0800 CST
*/
func ToZeroAM(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
}
