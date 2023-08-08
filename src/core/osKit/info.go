package osKit

import "runtime"

// GetGoroutineCount 当前进程中的协程数.
var GetGoroutineCount func() int = runtime.NumGoroutine
