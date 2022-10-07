package boolKit

type (
	Boolean struct {
		value bool
	}
)

func (c *Boolean) GetValue() bool {
	if c == nil {
		return false
	}
	return c.value
}

func NewBoolean(value bool) *Boolean {
	return &Boolean{
		value: value,
	}
}
