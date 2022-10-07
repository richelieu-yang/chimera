package floatKit

type (
	Float32 struct {
		value float32
	}

	Float64 struct {
		value float64
	}
)

func (c *Float32) GetValue() float32 {
	if c == nil {
		return 0
	}
	return c.value
}

func (c *Float64) GetValue() float64 {
	if c == nil {
		return 0
	}
	return c.value
}

func NewFloat32(value float32) *Float32 {
	return &Float32{
		value: value,
	}
}

func NewFloat64(value float64) *Float64 {
	return &Float64{
		value: value,
	}
}
