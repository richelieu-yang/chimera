package floatKit

import "github.com/shopspring/decimal"

func init() {
	// 除不尽时，保留的小数位数（默认为16）
	decimal.DivisionPrecision = 16
}
