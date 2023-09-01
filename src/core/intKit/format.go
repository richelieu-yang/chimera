package intKit

import "strconv"

var (
	IntToString func(i int) string = strconv.Itoa

	// FormatInt int64 => string
	/*
		@param base 进制，范围: [2, 36]
	*/
	FormatInt func(i int64, base int) string = strconv.FormatInt

	// FormatUint uint64 => string
	/*
		@param base 进制，范围: [2, 36]
	*/
	FormatUint func(i uint64, base int) string = strconv.FormatUint
)

// Int64ToString 类型转换: int64 => string
func Int64ToString(i int64) string {
	return strconv.FormatInt(i, 10)
}
