package strKit

// GetStringPointer 类型转换: string => *string
func GetStringPointer(s string) *string {
	return &s
}
