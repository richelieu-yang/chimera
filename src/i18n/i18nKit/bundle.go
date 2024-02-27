package i18nKit

import (
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"github.com/richelieu-yang/chimera/v3/src/config/propertiesKit"
	"github.com/richelieu-yang/chimera/v3/src/config/tomlKit"
	"github.com/richelieu-yang/chimera/v3/src/config/yaml/yamlKit"
	"github.com/richelieu-yang/chimera/v3/src/serialize/json/jsonKit"
	"golang.org/x/text/language"
)

// NewBundle
/*
PS:
(1) 返回的 *i18n.Bundle 实例，后续可以通过 Bundle.LoadMessageFile() 或 Bundle.ParseMessageFileBytes() 继续加载message file;
(2) Bundle.ParseMessageFileBytes(buf []byte, path string)，传参path可以是 文件名 || 文件路径（相对或绝对，可以不存在）.

@param defaultLanguage e.g. language.English || language.Chinese || language.SimplifiedChinese || language.TraditionalChinese
@param messageFilePaths (1) 可以为nil
						(2) 支持的文件类型: toml、json、yaml、properties
*/
func NewBundle(defaultLanguage language.Tag, messageFilePaths ...string) (*i18n.Bundle, error) {
	bundle := i18n.NewBundle(defaultLanguage)
	bundle.RegisterUnmarshalFunc("toml", tomlKit.Unmarshal)
	bundle.RegisterUnmarshalFunc("json", jsonKit.Unmarshal)
	bundle.RegisterUnmarshalFunc("yaml", yamlKit.Unmarshal)
	bundle.RegisterUnmarshalFunc("properties", propertiesKit.Unmarshal)

	for _, path := range messageFilePaths {
		if err := LoadMessageFile(bundle, path); err != nil {
			return nil, err
		}
	}
	return bundle, nil
}
