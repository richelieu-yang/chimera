package osKit

import (
	"runtime"
	"strconv"
)

// OS 操作系统
var OS string

// ARCH 处理器架构
var ARCH string

// BITS 操作系统的位数（32 || 64）
/*
参考: Go获取操作系统位数 https://blog.csdn.net/TCatTime/article/details/106815724
*/
var BITS int

func init() {
	OS = runtime.GOOS
	ARCH = runtime.GOARCH

	BITS = strconv.IntSize
}

func IsWindows() bool {
	return OS == `windows`
}

func IsMac() bool {
	return OS == `darwin`
}

func IsLinux() bool {
	return OS == `linux`
}
