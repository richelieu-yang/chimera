package stackTraceKit

import (
	"github.com/apache/rocketmq-clients/golang/v5/pkg/utils"
	"runtime/debug"
)

// PrintStackTrace 输出 stack trace 到 标准错误（os.Stderr）
/*
PS: 可参考《Golang.wps》.
*/
var PrintStackTrace func() = debug.PrintStack

var GetStackTrace func() []byte = debug.Stack

var GetAllStackTraces func() string = utils.DumpStacks
