package randomKit

import "github.com/duke-git/lancet/v2/random"

var (
	// RandFloat 生成随机float64数字，可以指定范围和精度.
	/*
		@param precision 精度（小数点后保留几位）
	*/
	RandFloat func(min, max float64, precision int) float64 = random.RandFloat

	// RandFloats 生成随机float64数字切片，指定长度，范围和精度.
	/*
		@param precision 精度（小数点后保留几位）
	*/
	RandFloats func(n int, min, max float64, precision int) []float64 = random.RandFloats
)
