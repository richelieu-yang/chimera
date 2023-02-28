package strKit

func GetStringPtr(str string) *string {
	return &str
}

// GetStringPtrValue
/*
@param ptr 不能为nil（会导致panic: runtime error: invalid memory address or nil pointer dereference）
*/
func GetStringPtrValue(ptr *string) string {
	return *ptr
}
