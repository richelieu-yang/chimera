package gcacheKit

import "github.com/gogf/gf/v2/os/gcache"

var (
	// New
	/*
		PS: 不使用时，需要手动调用 Close().
	*/
	New func(lruCap ...int) *gcache.Cache = gcache.New
)
