package jsonKit

// Marshal 序列化.
/*
@param obj 可以为nil || ""

e.g.
(nil) 	=> []byte("null"), nil
("") 	=> []byte("\"\""), nil
*/
func Marshal(obj interface{}, options ...JsonOption) ([]byte, error) {
	opts := loadOptions(options...)
	return opts.api.Marshal(obj)
}

func MarshalWithIndent(obj interface{}, options ...JsonOption) ([]byte, error) {
	prefix := ""
	// Richelieu: 与golang自带的"encoding/json"（可以用"\t"）不同，indent中不能有非空格的字符
	//indent := "\t"
	indent := "    "

	opts := loadOptions(options...)
	return opts.api.MarshalIndent(obj, prefix, indent)
}

// MarshalToString 序列化为字符串.
/*
@param obj 可以为nil || ""

e.g.
(nil) 	=> "null", nil
("") 	=> "\"\"", nil
*/
func MarshalToString(obj interface{}, options ...JsonOption) (string, error) {
	opts := loadOptions(options...)
	return opts.api.MarshalToString(obj)
}

func MarshalToStringWithIndent(obj interface{}, options ...JsonOption) (string, error) {
	data, err := MarshalWithIndent(obj, options...)
	if err != nil {
		return "", err
	}
	return string(data), nil
}
