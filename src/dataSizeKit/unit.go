package dataSizeKit

import (
	"github.com/shopspring/decimal"
)

const (
	B   uint64 = 1
	KiB uint64 = 1 << 10
	MiB uint64 = 1 << 20
	GiB uint64 = 1 << 30
	TiB uint64 = 1 << 40
	PiB uint64 = 1 << 50
	EiB uint64 = 1 << 60
	//ZB  uint64 = 1 << 70
	//YB  uint64 = 1 << 80
)

// ByteToKiB
/*
	PS: IEC标准.

	@param places 保留的小数位
*/
func ByteToKiB(bytes uint64, places int32) float64 {
	return div(bytes, KiB, places)
}

// ByteToMiB
/*
	PS: IEC标准.

	@param places 保留的小数位
*/
func ByteToMiB(bytes uint64, places int32) float64 {
	return div(bytes, MiB, places)
}

// ByteToGiB
/*
	PS: IEC标准.

	@param places 保留的小数位
*/
func ByteToGiB(bytes uint64, places int32) float64 {
	return div(bytes, GiB, places)
}

func div(bytes, uint uint64, placesArgs ...int32) float64 {
	var places int32
	if len(placesArgs) > 0 {
		places = placesArgs[0]
	} else {
		// 默认保留 2 位小数
		places = 2
	}

	d := decimal.NewFromInt(int64(bytes)).Div(decimal.NewFromInt(int64(uint))).Round(places)
	f, _ := d.Float64()
	return f
}
