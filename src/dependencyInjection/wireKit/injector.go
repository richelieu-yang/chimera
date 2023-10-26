package wireKit

import "github.com/google/wire"

var (
	// Build 声明一个注入器函数.
	/*
		PS:
		(1) 与提供者一样，注入器也可以输入参数（然后将其发送给提供者），并且可以返回错误.
		(2) wire.Build 的参数和 wire.NewSet 一样：都是提供者集合。这些就在该注入器的代码生成期间使用的提供者集。

		@return 这个函数的返回值也无关紧要，只要它们的类型正确即可，这些值在生成的代码中将被忽略
	*/
	Build func(...interface{}) string = wire.Build
)
