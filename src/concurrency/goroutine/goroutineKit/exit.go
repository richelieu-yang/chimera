package goroutineKit

import "runtime"

// Goexit 退出当前goroutine.
var Goexit func() = runtime.Goexit
