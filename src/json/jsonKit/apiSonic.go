// 编译标签(build tag): 参考gin中的"sonic.go"
//go:build sonic && avx && go1.16 && amd64 && (linux || windows || darwin)

package jsonKit

import (
	"github.com/bytedance/sonic"
	"github.com/klauspost/cpuid/v2"
	"github.com/richelieu-yang/chimera/v2/src/consts"
	"github.com/richelieu-yang/chimera/v2/src/core/cpuKit"
	"github.com/sirupsen/logrus"
)

/*
!!!: 并非 amd64 CPU 就行了，还需要支持 avx指令集 等.（e.g.yozo某台amd64内网机就不行）
*/
func init() {
	library = "bytedance/sonic"
	defaultApi = sonic.ConfigDefault
	stdApi = sonic.ConfigStd

	if !cpuKit.HasFeature(cpuid.AVX) {
		logrus.Fatalf("[%s, JSON] AVX isn't supported by CPU!!!", consts.UpperProjectName)
		return
	}

	testAPI()
}
