package funcKit

import (
	"fmt"
	"runtime"
	"strings"
)

// GetEntireCaller
/*
e.g.
(1) => test/test1.go:26 testFunc()
*/
func GetEntireCaller(callDepth int) string {
	pc, file, line, ok := runtime.Caller(callDepth)
	if !ok {
		return ""
	}
	return fmt.Sprintf("%s %s()", prettyCaller(file, line), prettyFuncName(pc))
}

func prettyFuncName(pc uintptr) string {
	// e.g."github.com/richelieu42/chimera/src/core/file/fileKit.AssertExistAndIsFile"
	funcName := runtime.FuncForPC(pc).Name()

	index := strings.LastIndex(funcName, ".")
	if index != -1 {
		funcName = funcName[index+1:]
	}
	return funcName
}

// GetCaller
/*
参考: go-zero中的logx/util.go.

@param callDepth 必须满足: >=0，实际使用中: >=1

e.g.
(1)	=> "test/test1.go:27"
*/
func GetCaller(callDepth int) string {
	_, file, line, ok := runtime.Caller(callDepth)
	if !ok {
		return ""
	}
	return prettyCaller(file, line)
}

func prettyCaller(file string, line int) string {
	idx := strings.LastIndexByte(file, '/')
	if idx < 0 {
		return fmt.Sprintf("%s:%d", file, line)
	}

	idx = strings.LastIndexByte(file[:idx], '/')
	if idx < 0 {
		return fmt.Sprintf("%s:%d", file, line)
	}

	return fmt.Sprintf("%s:%d", file[idx+1:], line)
}
