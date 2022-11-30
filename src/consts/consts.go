// Package consts
/*
本库的全局常量.
*/
package consts

import "github.com/richelieu42/go-scales/src/core/osKit"

// LeftParenthesisInFileName 左侧的小括号
var LeftParenthesisInFileName string

// RightParenthesisInFileName 右侧的小括号
var RightParenthesisInFileName string

func init() {
	if osKit.IsLinux() {
		// Linux: 半角
		LeftParenthesisInFileName = "("
		RightParenthesisInFileName = ")"
	} else {
		// Windows || Mac: 全角
		LeftParenthesisInFileName = "（"
		RightParenthesisInFileName = "）"
	}
}

const (
	// Name 本库（项目）的名字
	Name = "go-scales"

	EnvPath = "scales-config/env.yaml"

	Ip2RegionXdb = "scales-config/ip2region.xdb"
)
