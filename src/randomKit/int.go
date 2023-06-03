package randomKit

import "github.com/gogf/gf/v2/util/grand"

var (
	// Intn 随机生成int类型的值.
	/*
		@return 范围: [0, max)
	*/
	Intn func(max int) int = grand.Intn

	// Int 随机生成int类型的值.
	/*
	   @return 范围: [min, max]
	*/
	Int func(min, max int) int = grand.N
)
