package intKit

func NewIntPointer(i int) *int {
	return &i
}

func NewInt8Pointer(i int8) *int8 {
	return &i
}

func NewInt16Pointer(i int16) *int16 {
	return &i
}

func NewInt32Pointer(i int32) *int32 {
	return &i
}

func NewInt64Pointer(i int64) *int64 {
	return &i
}

func NewUint8Pointer(i uint8) *uint8 {
	return &i
}

func NewUint16Pointer(i uint16) *uint16 {
	return &i
}

func NewUint32Pointer(i uint32) *uint32 {
	return &i
}

func NewUint64Pointer(i uint64) *uint64 {
	return &i
}

//type (
//	// UnsignedInteger 无符号整数
//	UnsignedInteger struct {
//		value uint
//	}
//
//	Integer struct {
//		value int
//	}
//
//	Integer32 struct {
//		value int32
//	}
//
//	Integer64 struct {
//		value int64
//	}
//)
//
//func (c *UnsignedInteger) GetValue() uint {
//	if c == nil {
//		return 0
//	}
//	return c.value
//}
//
//func (c *Integer) GetValue() int {
//	if c == nil {
//		return 0
//	}
//	return c.value
//}
//
//func (c *Integer32) GetValue() int32 {
//	if c == nil {
//		return 0
//	}
//	return c.value
//}
//
//func (c *Integer64) GetValue() int64 {
//	if c == nil {
//		return 0
//	}
//	return c.value
//}
//
//func NewUnsignedInteger(value uint) *UnsignedInteger {
//	return &UnsignedInteger{
//		value: value,
//	}
//}
//
//func NewInteger(value int) *Integer {
//	return &Integer{
//		value: value,
//	}
//}
//
//func NewInteger32(value int32) *Integer32 {
//	return &Integer32{
//		value: value,
//	}
//}
//
//func NewInteger64(value int64) *Integer64 {
//	return &Integer64{
//		value: value,
//	}
//}
