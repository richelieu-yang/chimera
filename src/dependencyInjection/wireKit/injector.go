package wireKit

import "github.com/google/wire"

var (
	// Build 声明一个注入器函数.
	/*
		@return 这个函数的返回值也无关紧要，只要它们的类型正确即可，这些值在生成的代码中将被忽略
	*/
	Build func(...interface{}) string = wire.Build
)
