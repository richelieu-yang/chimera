package validateKit

// Struct 验证结构体.
/*
@param s 如果为nil，将返回error(e.g. validator: (nil *main.User))
*/
func Struct(s interface{}) error {
	if validatable, ok := s.(Validatable); ok {
		return validatable.Validate()
	}

	v := New()
	return v.Struct(s)
}
