package uintptrKit

import (
	"unsafe"
)

// IntToUintptr int => uintptr
func IntToUintptr(i int) uintptr {
	return uintptr(i)
}

// StringToUintptr string => uintptr
/*
参考：https://mp.weixin.qq.com/s/z4dz5HBzliIFetOzLrFMJg
*/
func StringToUintptr(str string) uintptr {
	return uintptr(unsafe.Pointer(&str))
	//return uintptr(unsafe.Pointer(syscall.StringBytePtr(str)))
}
