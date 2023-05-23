package funcKit

import (
	"runtime"
	"strings"
)

// GetFuncName
/*
@param callDepth 	(1) 必须满足: >=0，实际使用中: >=1
					(2) 如果callDepth == 1，则返回 调用此函数的函数的函数名（不含包名）
*/
func GetFuncName(callDepth int) string {
	pc, _, _, ok := runtime.Caller(callDepth)
	if !ok {
		return ""
	}
	name := runtime.FuncForPC(pc).Name()
	index := strings.LastIndex(name, ".")
	if index != -1 {
		name = name[index+1:]
	}
	return name
}
