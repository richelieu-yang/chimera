package yamlKit

import (
	"github.com/richelieu-yang/chimera/v2/src/core/fileKit"
	"gopkg.in/yaml.v3"
)

var Marshal func(in interface{}) (out []byte, err error) = yaml.Marshal

// MarshalToString
/*
@param in 建议为结构体实例指针 || map实例 || slice实例
*/
func MarshalToString(in interface{}) (string, error) {
	data, err := Marshal(in)
	return string(data), err
}

// MarshalToFile
/*
PS: 对 传参filePath 的验证和断言在 fileKit.WriteToFile 里面.

@param in		建议为结构体实例指针 || map实例 || slice实例
@param filePath (1) .yaml格式的文件
				(2) 不存在的话，会创建一个新的文件
				(3) 存在且是个文件的话，会 "覆盖" 掉旧的（并不会加到该文件的最后面）
*/
func MarshalToFile(in interface{}, filePath string) error {
	data, err := Marshal(in)
	if err != nil {
		return err
	}
	return fileKit.WriteToFile(data, filePath)
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
