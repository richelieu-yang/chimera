package yamlKit

import (
	"github.com/richelieu-yang/chimera/v2/src/core/fileKit"
	"github.com/richelieu-yang/chimera/v2/src/core/ioKit"
	"github.com/richelieu-yang/chimera/v2/src/json/jsonKit"
	"github.com/spf13/viper"
)

// MarshalToFileWithJsonTag
/*
适用场景: 有json tag，没有yaml tag.
缺陷: 传参in 不能为slice实例.
*/
func MarshalToFileWithJsonTag(in interface{}, filePath string) error {
	if err := fileKit.AssertNotExistOrIsFile(filePath); err != nil {
		return err
	}
	if err := fileKit.MkParentDirs(filePath); err != nil {
		return err
	}

	// 结构体实例 => json
	data, err := jsonKit.Marshal(in)
	if err != nil {
		return err
	}

	// viper读取json，然后导出为yaml
	v := viper.New()
	v.SetConfigType("json")
	if err := v.ReadConfig(ioKit.NewReader(data)); err != nil {
		return err
	}
	return v.WriteConfigAs(filePath)
}

// MarshalToFileWithJsonTag1
/*
适用场景: 有json tag，没有yaml tag.
缺陷: 传参in 不能为slice实例.
*/
func MarshalToFileWithJsonTag1(in interface{}, filePath string) error {
	// 结构体实例 => json
	jsonData, err := jsonKit.Marshal(in)
	if err != nil {
		return err
	}

	// json => map
	m := map[string]interface{}{}
	if err := jsonKit.Unmarshal(jsonData, &m); err != nil {
		return err
	}

	// map => yaml文本
	return MarshalToFile(m, filePath)
}
