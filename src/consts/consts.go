// Package consts
/*
本库的全局常量.
*/
package consts

import "github.com/richelieu42/chimera/src/core/osKit"

// ColonInFileName （文件名中的）冒号
var ColonInFileName string

func init() {
	if osKit.IsLinux() {
		// Linux: 半角（英文）
		ColonInFileName = ":"
	} else {
		// Windows || Mac: 全角（中文）
		ColonInFileName = "："
	}
}

const (
	// Name 本库（项目）的名字
	Name = "go-scales"

	EnvPath = "scales-config/env.yaml"

	Ip2RegionXdb = "scales-config/ip2region.xdb"
)
