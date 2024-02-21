package viperKit

func UnmarshalProperties(content []byte, configType string, defaultMap map[string]interface{}, ptr interface{}) error {
	_, err := Unmarshal(content, "properties", nil, ptr)
	return err
}
