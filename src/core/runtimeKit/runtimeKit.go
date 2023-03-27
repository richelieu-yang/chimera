// Package runtimeKit
/**
 * 主要是对如下包的封装："runtime"
 */
package runtimeKit

import (
	"github.com/shirou/gopsutil/v3/host"
	"runtime"
)

// GoVersion Golang的版本号
var GoVersion string

// GoRoot GOROOT环境变量
var GoRoot string

func init() {
	GoVersion = runtime.Version()
	GoRoot = runtime.GOROOT()
}

func GetHostInfo() (*host.InfoStat, error) {
	return host.Info()
}
