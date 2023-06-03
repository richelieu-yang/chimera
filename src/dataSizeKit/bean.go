package dataSizeKit

import (
	"github.com/richelieu-yang/chimera/v2/src/core/floatKit"
	"github.com/richelieu-yang/chimera/v2/src/core/sliceKit"
	"math"
)

type (
	DataSize struct {
		// 数值
		Number float64
		// 单位
		Unit *Unit
	}
)

// ConvertToTargetUint 单位转换
/*
@param targetUnit 可以为nil，将使用默认单位
*/
func (size *DataSize) ConvertToTargetUint(targetUnit *Unit) *DataSize {
	if size.Unit.GetValue() == targetUnit.GetValue() {
		return size
	}

	size.Number = floatKit.Div(floatKit.Mul(size.Number, float64(size.Unit.GetValue())), float64(targetUnit.GetValue()))
	size.Unit = targetUnit
	return size
}

// GetByteValue
/*
@return 单位: B(byte)
*/
func (size *DataSize) GetByteValue() uint64 {
	tmp := floatKit.Mul(size.Number, float64(size.Unit.GetValue()))
	tmp = math.Ceil(tmp)
	return uint64(tmp)
}

func (size *DataSize) ToSuitableUint() *DataSize {
	bytes := size.GetByteValue()

	var unit *Unit
	if bytes < KB.GetValue() {
		unit = B
	} else if bytes < MB.GetValue() {
		unit = KB
	} else if bytes < GB.GetValue() {
		unit = MB
	} else if bytes < TB.GetValue() {
		unit = GB
	} else {
		unit = TB
	}

	size.ConvertToTargetUint(unit)
	return size
}

// ToString
/*
@param args 默认（不传参）不对小数位进行处理
*/
func (size *DataSize) ToString(precArgs ...int) string {
	// prec默认值: -1（代表使用最少数量的、但又必需的数字来表示f）
	prec := sliceKit.GetFirstItemWithDefault(-1, precArgs...)

	return floatKit.FormatFloat64ToString(size.Number, 'f', prec) + " " + size.Unit.String()
}
