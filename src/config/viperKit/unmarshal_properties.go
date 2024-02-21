package viperKit

func UnmarshalProperties(content []byte, ptr interface{}) error {
	_, err := Unmarshal(content, "properties", nil, ptr)
	return err
}
