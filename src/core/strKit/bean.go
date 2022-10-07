package strKit

type (
	String struct {
		value string
	}
)

func (c *String) GetValue() string {
	if c == nil {
		return ""
	}
	return c.value
}

func NewString(value string) *String {
	return &String{
		value: value,
	}
}
