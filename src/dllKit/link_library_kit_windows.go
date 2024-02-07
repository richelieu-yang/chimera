package linkLibraryKit

import (
	"github.com/richelieu-yang/chimera/v3/src/core/errorKit"
	"plugin"
)

// LoadLinkLibrary 加载动态链接库（Linux、Mac）
/*
TODO: 看后续"plugin标准库"是否会支持Windows环境.
*/
func LoadLinkLibrary(path string) (*plugin.Plugin, error) {
	return nil, errorKit.New("Link libraries cannot be loaded in Windows!")
}
