package linkLibraryKit

import (
	"github.com/richelieu42/go-scales/src/core/errorKit"
	"github.com/richelieu42/go-scales/src/core/file/fileKit"
	"github.com/richelieu42/go-scales/src/core/strKit"
	"plugin"
)

// LoadLinkLibrary 加载动态链接库（Linux、Mac）
func LoadLinkLibrary(path string) (*plugin.Plugin, error) {
	if strKit.IsEmpty(path) {
		return nil, errorKit.Simple("path of link library is empty")
	}
	if !fileKit.Exist(path) {
		return nil, errorKit.Simple("link library(path: %s) doesn't exist", path)
	}
	return plugin.Open(path)
}
