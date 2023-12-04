package viperKit

import (
	"github.com/richelieu-yang/chimera/v2/src/core/sliceKit"
	"github.com/richelieu-yang/chimera/v2/src/core/strKit"
	"github.com/richelieu-yang/chimera/v2/src/file/fileKit"
	"github.com/spf13/viper"
)

// IsContentTypeSupported
/*
@param extName "json", "toml", "yaml", "yml", "properties", "props", "prop", "hcl", "tfvars", "dotenv", "env", "ini"
*/
func IsContentTypeSupported(contentType string) bool {
	return sliceKit.Contains(viper.SupportedExts, contentType)
}

// PolyfillContentType
/*
e.g.
	extname := ".JSON"
	fmt.Println(viperKit.PolyfillContentType(extname)) // "json"

@param extName "json", "toml", "yaml", "yml", "properties", "props", "prop", "hcl", "tfvars", "dotenv", "env", "ini"
*/
func PolyfillContentType(contentType string) string {
	contentType = strKit.TrimSpace(contentType)
	contentType = strKit.ToLower(contentType)

	if len(contentType) > 1 && strKit.StartWith(contentType, ".") {
		contentType = strKit.RemovePrefixIfExists(contentType, ".")
	}
	return contentType
}

func GetContentType(path string) string {
	contentType := fileKit.GetExtName(path)
	return PolyfillContentType(contentType)
}
