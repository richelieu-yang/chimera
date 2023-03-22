package timeKit

import (
	"time"
)

type (
	TimeFormat string
)

const (
	Day  = 24 * time.Hour
	Week = 7 * Day
)

/*
UTC: 世界协调时间(Coordinated Universal Time)
GMT: 格林威治时间(UTC+0)
CST: 中国标准时间(UTC+8)
MST: 北美山区标准时间(UTC-7)
*/
const (
	// FormatDefault 参考"format.go"
	FormatDefault TimeFormat = "2006-01-02 15:04:05.999999999 -0700 MST"

	// FormatFileName 用于作为文件名（或目录名）的一部分
	FormatFileName TimeFormat = "2006-01-02T15.04.05.000"

	// FormatCommon 常规的格式
	FormatCommon TimeFormat = "2006-01-02T15:04:05.000"
	// FormatCommon1 常规的格式1
	FormatCommon1 TimeFormat = "2006-01-02 15:04:05.000"

	// FormatEntire 完整的格式
	FormatEntire TimeFormat = "2006-01-02 15:04:05.000Z07:00"

	FormatA TimeFormat = "2006-01-02 15:04:05"
	FormatB TimeFormat = "2006-01-02 3:04:05.000 PM Mon Jan"
	FormatC TimeFormat = "3:04:05.000 PM Mon Jan"

	// FormatNetwork 网络的格式
	FormatNetwork TimeFormat = "Mon, 02 Jan 2006 15:04:05 MST"
)
