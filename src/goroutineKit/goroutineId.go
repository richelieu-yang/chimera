package goroutineKit

import (
	"bytes"
	"runtime"
	"strconv"
)

// GetGoroutineID 获取协程的id
// DEPRECATED: 在go语言中，谷歌开发者不建议大家获取协程ID，主要是为了GC更好的工作，滥用协程ID会导致GC不能及时回收内存。如果一个第三方库使用了协程ID，那么使用该库的人将会莫名中招。
/*
golang 获取goroutineID
	https://studygolang.com/articles/19794
Golang - 获取协程ID，从此走上一条曲折的不归路
	https://blog.csdn.net/wyansai/article/details/101004226
*/
func GetGoroutineID() uint64 {
	b := make([]byte, 64)
	runtime.Stack(b, false)
	b = bytes.TrimPrefix(b, []byte("goroutine "))
	b = b[:bytes.IndexByte(b, ' ')]
	n, _ := strconv.ParseUint(string(b), 10, 64)
	return n
}
