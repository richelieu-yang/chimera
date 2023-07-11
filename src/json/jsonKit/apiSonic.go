// 编译标签(build tag): 参考gin中的"sonic.go"
//go:build (linux || windows || darwin) && amd64

package jsonKit

import (
	"github.com/bytedance/sonic"
)

func init() {
	api = sonic.ConfigDefault
}
