package confKit

import "github.com/zeromicro/go-zero/core/conf"

// LoadFromYamlBytes 加载 .yaml 格式的配置文件内容.
var LoadFromYamlBytes func(content []byte, v any) error = conf.LoadFromYamlBytes

func LoadFromYamlText(text string, v any) error {
	return LoadFromYamlBytes([]byte(text), v)
}
