package confKit

import "github.com/zeromicro/go-zero/core/conf"

// LoadFromTomlBytes 加载 .toml 格式的配置文件内容.
var LoadFromTomlBytes func(content []byte, v any) error = conf.LoadFromTomlBytes

func LoadFromTomlText(text string, v any) error {
	return LoadFromTomlBytes([]byte(text), v)
}
