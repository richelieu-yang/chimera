// Package runtimeKit
/**
 * 主要是对如下包的封装："runtime"
 */
package runtimeKit

import (
	"github.com/richelieu42/chimera/src/core/errorKit"
	"github.com/shirou/gopsutil/v3/host"
	"runtime"
)

// GoVersion Golang的版本号
var GoVersion string

// GoRoot GOROOT环境变量
var GoRoot string

var hostInfo *host.InfoStat

func init() {
	GoVersion = runtime.Version()
	GoRoot = runtime.GOROOT()

	var err error
	hostInfo, err = host.Info()
	if err != nil {
		errorKit.Panic("fail to get host info, error:\n%+v", err)
	}
}

func GetHostInfo() *host.InfoStat {
	return hostInfo
}
