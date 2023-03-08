package dataSizeKit

import (
	"github.com/richelieu42/go-scales/src/core/errorKit"
	"github.com/richelieu42/go-scales/src/core/floatKit"
	"github.com/richelieu42/go-scales/src/core/strKit"
)

// ToReadableDataSizeString
/*
e.g.
(1725 * 1024) => "1.68 MB"
*/
func ToReadableDataSizeString(bytes uint64) string {
	size := &DataSize{
		Number: float64(bytes),
		Unit:   B,
	}
	return size.ToSuitableUint().ToString(2)
}

// ParseStringToDataSize string => *DataSize
func ParseStringToDataSize(sizeStr string) (*DataSize, error) {
	str := strKit.RemoveSpace(sizeStr)
	if strKit.IsEmpty(str) {
		return nil, errorKit.Simple("sizeStr(%s) is invalid", sizeStr)
	}
	str = strKit.ToUpper(str)
	str = strKit.AppendIfMissing(str, "B")

	var unit *Unit
	if strKit.EndWith(str, TB.String()) {
		unit = TB
	} else if strKit.EndWith(str, GB.String()) {
		unit = GB
	} else if strKit.EndWith(str, MB.String()) {
		unit = MB
	} else if strKit.EndWith(str, KB.String()) {
		unit = KB
	} else if strKit.EndWith(str, B.String()) {
		unit = B
	} else {
		// 理论上，代码不会走到此处
		return nil, errorKit.Simple("sizeStr(%s) is invalid", sizeStr)
	}

	numberStr := strKit.RemoveSuffixIfExist(str, unit.String())
	number, err := floatKit.ParseStringToFloat64(numberStr)
	if err != nil || number <= 0 {
		return nil, errorKit.Simple("sizeStr(%s) is invalid", sizeStr)
	}

	return &DataSize{
		Number: number,
		Unit:   unit,
	}, nil
}
