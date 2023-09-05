package yamlKit

import "gopkg.in/yaml.v3"

var Marshal func(in interface{}) (out []byte, err error) = yaml.Marshal

// MarshalToString
/*
@param in 建议为结构体实例指针 || map实例 || slice实例
*/
func MarshalToString(in interface{}) (string, error) {
	data, err := Marshal(in)
	return string(data), err
}

// Unmarshal
/*
Deprecated: Use confKit.MustLoad || confKit.LoadFromYamlBytes instead.
*/
var Unmarshal func(in []byte, out interface{}) (err error) = yaml.Unmarshal

// UnmarshalFromString
/*
Deprecated: Use confKit.MustLoad || confKit.LoadFromYamlBytes instead.
*/
func UnmarshalFromString(in string, out interface{}) error {
	return Unmarshal([]byte(in), out)
}
