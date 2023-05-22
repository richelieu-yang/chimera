package funcKit

import (
	"github.com/richelieu42/chimera/v2/src/core/strKit"
	"runtime"
)

// GetFuncName
/*
@param callDepth 必须满足: >=0，实际使用中: >=1

e.g.
(1) => 调用此函数的函数的函数名（不含包名）
*/
func GetFuncName(callDepth int) string {
	pc, _, _, ok := runtime.Caller(callDepth)
	if !ok {
		return ""
	}
	name := runtime.FuncForPC(pc).Name()
	index := strKit.LastIndex(name, ".")
	if index != -1 {
		name = strKit.SubAfter(name, index+1)
	}
	return name
}
