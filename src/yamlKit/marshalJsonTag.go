package yamlKit

import (
	"github.com/richelieu-yang/chimera/v2/src/json/jsonKit"
)

// MarshalToFileWithJsonTag
/*
适用场景: 有json tag，没有yaml tag.
*/
func MarshalToFileWithJsonTag(in interface{}, filePath string) error {
	// 结构体实例 => json
	jsonData, err := jsonKit.Marshal(in)
	if err != nil {
		return err
	}

	// json => map
	m := map[string]interface{}{}
	if err := jsonKit.Unmarshal(jsonData, m); err != nil {
		return err
	}

	// map => yaml文本
	return MarshalToFile(m, filePath)
}
