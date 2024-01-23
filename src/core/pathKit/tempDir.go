package pathKit

import (
	"github.com/richelieu-yang/chimera/v2/src/consts"
	"github.com/richelieu-yang/chimera/v2/src/core/strKit"
	"github.com/richelieu-yang/chimera/v2/src/file/fileKit"
	"os"
)

var tempDir string

// GetTempDir 获取系统的临时目录.
/*
Deprecated: 不建议向系统临时目录中放东西（服务器不一定会给权限）.

e.g.
	Windows	"C:\Users\Lenovo\AppData\Local\Temp"
	Mac		"/var/folders/4_/33p_vn057msfh2nvgx6hwv_40000gn/T/"
*/
var GetTempDir func() string = os.TempDir

// GetExclusiveTempDir 获取 本依赖 的专属临时目录.
/*
	PS: 不建议向系统临时目录中放东西（服务器不一定会给权限）.

	e.g. Mac
		() => "/var/folders/4_/33p_vn057msfh2nvgx6hwv_40000gn/T/$$chimera", nil
*/
func GetExclusiveTempDir() (string, error) {
	if strKit.IsNotEmpty(tempDir) {
		return tempDir, nil
	}

	dir := Join(GetTempDir(), "$$"+consts.ProjectName)
	if err := fileKit.MkDirs(dir); err != nil {
		return "", err
	}
	return dir, nil
}

func SetTempDir(dirPath string) error {
	if err := fileKit.AssertNotExistOrIsDir(dirPath, true); err != nil {
		return err
	}

	tempDir = dirPath
	return nil
}
