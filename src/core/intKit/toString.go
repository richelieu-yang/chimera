package intKit

import "strconv"

func IntToString(i int) string {
	return strconv.Itoa(i)
}

// FormatIntToString 类型转换: int => string
/*
@param base 进制（有效值: [2, 36]）
*/
func FormatIntToString(i int, base int) string {
	return FormatInt64ToString(int64(i), base)
}

// FormatInt8ToString 类型转换: int8 => string
/*
@param base 进制（有效值: [2, 36]）
*/
func FormatInt8ToString(i int8, base int) string {
	return FormatInt64ToString(int64(i), base)
}

// FormatInt16ToString 类型转换: int16 => string
/*
@param base 进制（有效值: [2, 36]）
*/
func FormatInt16ToString(i int16, base int) string {
	return FormatInt64ToString(int64(i), base)
}

// FormatInt32ToString 类型转换: int32 => string
/*
@param base 进制（有效值: [2, 36]）
*/
func FormatInt32ToString(i int32, base int) string {
	return FormatInt64ToString(int64(i), base)
}

// FormatInt64ToString 类型转换: int64 => string
/*
@param base 进制（有效值: [2, 36]）
*/
func FormatInt64ToString(i int64, base int) string {
	return strconv.FormatInt(i, base)
}

// FormatUint8ToString 类型转换: uint8 => string
/*
@param base 进制（有效值: [2, 36]）
*/
func FormatUint8ToString(i uint8, base int) string {
	return strconv.FormatUint(uint64(i), base)
}

// FormatUint16ToString 类型转换: uint16 => string
/*
@param base 进制（有效值: [2, 36]）
*/
func FormatUint16ToString(i uint16, base int) string {
	return strconv.FormatUint(uint64(i), base)
}

// FormatUint32ToString 类型转换: uint32 => string
/*
@param base 进制（有效值: [2, 36]）
*/
func FormatUint32ToString(i uint32, base int) string {
	return strconv.FormatUint(uint64(i), base)
}

// FormatUint64ToString 类型转换: uint64 => string
/*
@param base 进制（有效值: [2, 36]）
*/
func FormatUint64ToString(i uint64, base int) string {
	return strconv.FormatUint(i, base)
}
