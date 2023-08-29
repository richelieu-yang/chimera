package jsonKit

func Marshal(v interface{}) ([]byte, error) {
	return defaultApi.Marshal(v)
}

// MarshalIndent
/*
@param indent 为了兼容性，用"    "（4个空格）替代"\t"
*/
func MarshalIndent(v interface{}, prefix, indent string) ([]byte, error) {
	return defaultApi.MarshalIndent(v, prefix, indent)
}

func MarshalToString(v interface{}) (string, error) {
	return defaultApi.MarshalToString(v)
}

// MarshalIndentToString
/*
@param indent 为了兼容性，用"    "（4个空格）替代"\t"
*/
func MarshalIndentToString(v interface{}, prefix, indent string) (string, error) {
	data, err := MarshalIndent(v, prefix, indent)
	return string(data), err
}

func Unmarshal(data []byte, v interface{}) error {
	return defaultApi.Unmarshal(data, v)
}

func UnmarshalFromString(str string, v interface{}) error {
	return defaultApi.UnmarshalFromString(str, v)
}
