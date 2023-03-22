// Package runtimeKit
/**
 * 主要是对如下包的封装："runtime"
 */
package runtimeKit

import (
	"github.com/shirou/gopsutil/v3/host"
	"github.com/sirupsen/logrus"
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
		logrus.WithFields(logrus.Fields{
			"error": err.Error(),
		}).Fatal("fail to get host info")
	}
}

func GetHostInfo() *host.InfoStat {
	return hostInfo
}
