package funcKit

import (
	"github.com/richelieu42/chimera/src/consts/key"
	"runtime"
	"strings"
)

// GetCallerDetail 获取 调用此函数的函数 的详细信息（函数名、文件路径、行数）.
/*
PS: 由于此方法多了一层调用，所以需要+1，最终skip为2.

@return e.g. main.c /Users/richelieu/GolandProjects/go-scales/test/test1.go 32 true
*/
func GetCallerDetail() (string, string, int, bool) {
	return GetCallerDetailWithSkip(2)
}

// GetCallerDetailWithSkip
/*
参考:
	Golang实现获取当前函数名称和文件行号等操作 https://www.51sjk.com/b123b248673/
	Go学习——runtime.Caller()函数 https://blog.csdn.net/weixin_52000204/article/details/124504877

PS: 有可能获取失败，但目前不清楚什么情况下会失败（交叉编译时加上 -ldflags="-w -s" 并不会导致获取失败）.

@param skip 理论上此值应该>=1，1: 向上1层（即谁调用此方法的）；2: 向上2层...
@return e.g. ("main.main", "/Users/richelieu/GolandProjects/go-scales/test/test1.go", 10)
*/
func GetCallerDetailWithSkip(skip int) (funcName string, filePath string, line int, ok bool) {
	/*
		runtime.Caller的传参skip（常用的是1和2）: 上溯的栈帧数，0表示Caller的调用者（Caller所在的调用栈）（0-当前函数，1-上一层函数，…）.
		0		当前函数
		1		调用此函数的函数
		2		调用此函数的函数 * 2
		...
	*/
	pc, filePath, line, ok := runtime.Caller(skip)
	if ok {
		funcName = runtime.FuncForPC(pc).Name()
		funcName = processFuncName(funcName)
	} else {
		funcName = key.UNKNOWN
		filePath = key.UNKNOWN
		line = -1
	}
	return
}

// GetCallerName 获取 调用此函数的函数 的名称（上1层）.
func GetCallerName() string {
	return GetCallerNameWithSkip(2)
}

// GetCallerNameWithSkip
/*
@param skip 1: 向上1层
*/
func GetCallerNameWithSkip(skip int) string {
	// 此处调用多了1层，因此需要"+ 1"
	funcName, _, _, _ := GetCallerDetailWithSkip(skip + 1)
	return funcName
}

// processFuncName 处理特殊情况: 包名太长的情况
/*
e.g.
("github.com/richelieu42/chimera/src/core/file/fileKit.AssertExistAndIsFile") = > "fileKit.AssertExistAndIsFile"
*/
func processFuncName(name string) string {
	//if !strings.HasSuffix(name, "/") {
	index := strings.LastIndex(name, "/")
	if index != -1 {
		name = name[index+1:]
	}
	//}
	return name
}
