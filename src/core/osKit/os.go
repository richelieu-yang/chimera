package osKit

import (
	"github.com/dablelv/cyan/os"
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

var (
	IsWindows func() bool = os.IsWin

	IsLinux func() bool = os.IsLinux

	IsMac func() bool = os.IsMac
)
