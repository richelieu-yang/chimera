package bytesKit

import "unsafe"

// StringToBytes string => []byte
/*
PS:
(1) 参考了 gin(v1.8.1) 中的 internal/bytesconv/bytesconv.go ，后续要及时同步更新此方法！
(2) 必须满足的条件：对数据仅仅只有读操作！！！
*/
func StringToBytes(str string) []byte {
	return *(*[]byte)(unsafe.Pointer(
		&struct {
			string
			Cap int
		}{str, len(str)},
	))
}

// BytesToString []byte => string
/*
PS:
(1) 参考了 gin(v1.8.1) 中的 internal/bytesconv/bytesconv.go ，后续要及时同步更新此方法！
(2) 必须满足的条件：对数据仅仅只有读操作！！！
*/
func BytesToString(bytes []byte) string {
	return *(*string)(unsafe.Pointer(&bytes))
}
