package viperKit

import (
	"github.com/richelieu-yang/chimera/v2/src/core/sliceKit"
	"github.com/richelieu-yang/chimera/v2/src/core/strKit"
	"github.com/spf13/viper"
)

// IsExtNameSupported
/*
@param extName 	(1) <=> configType
				(2) "json", "toml", "yaml", "yml", "properties", "props", "prop", "hcl", "tfvars", "dotenv", "env", "ini"
*/
func IsExtNameSupported(extName string) bool {
	return sliceKit.Contains(viper.SupportedExts, extName)
}

// PolyfillExtName
/*
e.g.
	extname := ".JSON"
	fmt.Println(viperKit.PolyfillExtName(extname)) // "json"

@param extName 	(1) <=> configType
				(2) "json", "toml", "yaml", "yml", "properties", "props", "prop", "hcl", "tfvars", "dotenv", "env", "ini"
*/
func PolyfillExtName(extName string) string {
	extName = strKit.TrimSpace(extName)
	extName = strKit.ToLower(extName)

	if len(extName) > 1 && strKit.StartWith(extName, ".") {
		extName = strKit.RemovePrefixIfExists(extName, ".")
	}
	return extName
}
