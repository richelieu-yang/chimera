package viperKit

// UnmarshalProperties Deprecated: 反序列化后，key中的大写字母会被转换为小写字母.
func UnmarshalProperties(content []byte, ptr interface{}) error {
	_, err := Unmarshal(content, "properties", nil, ptr)
	return err
}
