package strKit

import (
	"unsafe"
)

// StringToBytes converts string to byte slice without a memory allocation.
/*
PS:
(1) copy from gin/internal/bytesconv/bytesconv.go
(2) Richelieu: 感觉这样效率更高，但慎用!!!
*/
func StringToBytes(s string) []byte {
	return *(*[]byte)(unsafe.Pointer(
		&struct {
			string
			Cap int
		}{s, len(s)},
	))
}

// BytesToString converts byte slice to string without a memory allocation.
/*
PS:
(1) copy from gin/internal/bytesconv/bytesconv.go
(2) Richelieu: 感觉这样效率更高，但慎用!!!
*/
func BytesToString(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}
