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
	// DirFormat 用于作为目录名（或者其中的一部分）
	DirFormat TimeFormat = "2006-01-02T15.04.05.000"

	// CommonFormat 常规的格式
	CommonFormat TimeFormat = "2006-01-02T15:04:05.000"
	// CommonFormat1 常规的格式1
	CommonFormat1 TimeFormat = "2006-01-02 15:04:05.000"

	// EntireFormat 完整的格式
	EntireFormat TimeFormat = "2006-01-02 15:04:05.000Z07:00"

	FormatA TimeFormat = "2006-01-02 15:04:05"
	FormatB TimeFormat = "2006-01-02 3:04:05.000 PM Mon Jan"
	FormatC TimeFormat = "3:04:05.000 PM Mon Jan"

	// NetworkFormat 网络的格式
	NetworkFormat TimeFormat = "Mon, 02 Jan 2006 15:04:05 MST"

	// DefaultFormat 参考"format.go"
	DefaultFormat TimeFormat = "2006-01-02 15:04:05.999999999 -0700 MST"
)
