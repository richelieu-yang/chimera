// 编译标签(build tag): 参考gin中的"sonic.go"
//go:build amd64 && go1.16

package jsonKit

import "github.com/bytedance/sonic"

/*
!!!: 并非 amd64 CPU 就行了，还需要支持 avx指令集 等.（e.g.yozo某台amd64内网机就不行）
*/
func init() {
	library = "bytedance/sonic"
	defaultAPI = sonic.ConfigDefault
}
