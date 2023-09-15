package confKit

import "github.com/zeromicro/go-zero/core/conf"

// LoadFromJsonBytes 加载 .json 格式的配置文件内容.
var LoadFromJsonBytes func(content []byte, v any) error = conf.LoadFromJsonBytes

func LoadFromJsonText(text string, v any) error {
	return LoadFromJsonBytes([]byte(text), v)
}
