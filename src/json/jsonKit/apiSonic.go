// 编译标签(build tag): 参考gin中的"sonic.go"
//go:build amd64 && go1.16

package jsonKit

import (
	"github.com/bytedance/sonic"
	"github.com/sirupsen/logrus"
)

/*
!!!: 并非 amd64 CPU 就行了，还需要支持 avx指令集 等.（e.g.yozo某台amd64内网机就不行）
*/
func init() {
	library = "bytedance/sonic"
	api = sonic.ConfigDefault

	/*
		amd64 CPU，不支持 avx指令集 的情况下，下面的代码会报错 SIGILL: illegal instruction
		（启动时报错 总好过 运行时报错）
	*/
	m := map[string]interface{}{
		"0": 3.1415926,
		"1": 1,
	}
	jsonStr, err := sonic.ConfigStd.MarshalToString(m)
	if err != nil {
		logrus.Fatal(err)
	}
	logrus.Infof("jsonStr: [%s].", jsonStr)
	var m1 map[string]interface{}
	if err := sonic.ConfigStd.UnmarshalFromString(jsonStr, &m1); err != nil {
		logrus.Fatal(err)
	}
}
