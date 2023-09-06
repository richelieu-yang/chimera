package yamlKit

import "gopkg.in/yaml.v3"

// Unmarshal
/*
Deprecated: Use confKit.MustLoad || confKit.LoadFromYamlBytes instead.

PS: 需要搭配 yaml tag 一起使用，不识别 json tag.
*/
var Unmarshal func(in []byte, out interface{}) (err error) = yaml.Unmarshal

// UnmarshalFromString
/*
Deprecated: Use confKit.MustLoad || confKit.LoadFromYamlBytes instead.

PS: 需要搭配 yaml tag 一起使用，不识别 json tag.
*/
func UnmarshalFromString(in string, out interface{}) error {
	return Unmarshal([]byte(in), out)
}
