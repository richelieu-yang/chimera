// 编译标签(build tag): 参考gin中的"sonic.go"
//go:build sonic && avx && go1.16 && amd64 && (linux || windows || darwin)

package jsonKit

import (
	"github.com/bytedance/sonic"
	"github.com/klauspost/cpuid/v2"
	"github.com/richelieu-yang/chimera/v2/src/compareKit"
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
		logrus.Fatalf("[%s, JSON] AVX isn't supported by CPU!!!", consts.ProjectUpperName)
		return
	}

	/*
		amd64 CPU && 不支持 avx指令集 的情况下，下面的代码会报错 SIGILL: illegal instruction
		PS: 启动时报错退出程序 总好过 运行时报错退出程序.
	*/
	api := sonic.ConfigStd
	m := map[string]interface{}{
		"0": 3.1415926,
		"1": 1,
		"2": true,
		"3": `~!@#$%^&*()_+{}|:"><?	
qwdqw强	\t\r\n无敌
qwdqwd
威尔法
496465~·《》？L:"{}|,./l;'[]\/*-的确
`,
	}
	jsonData, err := api.Marshal(m)
	if err != nil {
		logrus.WithError(err).Fatal("[%s, JSON] Sonic fails to marshal!!!", consts.ProjectUpperName)
	}
	var m1 map[string]interface{}
	if err := api.Unmarshal(jsonData, &m1); err != nil {
		logrus.WithError(err).Fatal("[%s, JSON] Sonic fails to unmarshal!!!", consts.ProjectUpperName)
	}
	if compareKit.Equal(m, m1) {
		logrus.Fatalf("[%s, JSON] m and m1 aren't identical!!!", consts.ProjectUpperName)
	}
}
