package linkLibraryKit

import (
	"github.com/richelieu-yang/chimera/v2/src/core/errorKit"
	"github.com/richelieu-yang/chimera/v2/src/core/fileKit"
	"github.com/richelieu-yang/chimera/v2/src/core/strKit"
	"plugin"
)

// LoadLinkLibrary 加载动态链接库（Linux、Mac）
func LoadLinkLibrary(path string) (*plugin.Plugin, error) {
	if strKit.IsEmpty(path) {
		return nil, errorKit.New("path of link library is empty")
	}
	if !fileKit.Exist(path) {
		return nil, errorKit.New("link library(path: %s) doesn't exist", path)
	}
	return plugin.Open(path)
}
