package intKit

type (
	// UnsignedInteger 无符号整数
	UnsignedInteger struct {
		value uint
	}

	Integer struct {
		value int
	}

	Integer32 struct {
		value int32
	}

	Integer64 struct {
		value int64
	}
)

func (c *UnsignedInteger) GetValue() uint {
	if c == nil {
		return 0
	}
	return c.value
}

func (c *Integer) GetValue() int {
	if c == nil {
		return 0
	}
	return c.value
}

func (c *Integer32) GetValue() int32 {
	if c == nil {
		return 0
	}
	return c.value
}

func (c *Integer64) GetValue() int64 {
	if c == nil {
		return 0
	}
	return c.value
}

func NewUnsignedInteger(value uint) *UnsignedInteger {
	return &UnsignedInteger{
		value: value,
	}
}

func NewInteger(value int) *Integer {
	return &Integer{
		value: value,
	}
}

func NewInteger32(value int32) *Integer32 {
	return &Integer32{
		value: value,
	}
}

func NewInteger64(value int64) *Integer64 {
	return &Integer64{
		value: value,
	}
}
