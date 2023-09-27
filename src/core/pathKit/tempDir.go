package pathKit

import (
	"github.com/richelieu-yang/chimera/v2/src/consts"
	"github.com/richelieu-yang/chimera/v2/src/core/fileKit"
	"github.com/richelieu-yang/chimera/v2/src/core/strKit"
	"os"
)

var tempDir string

// GetOsTempDir 获取系统的临时目录.
/*
Deprecated: 不建议向系统临时目录中放东西（服务器不一定会给权限）.

e.g.
Windows: 	"C:\Users\Lenovo\AppData\Local\Temp"
Mac: 		"/var/folders/4_/33p_vn057msfh2nvgx6hwv_40000gn/T/"
*/
var GetOsTempDir func() string = os.TempDir

// GetTempDir 获取 本依赖 的专属临时目录.
/*
	PS: 不建议向系统临时目录中放东西（服务器不一定会给权限）.

	e.g. Mac
		() => "/var/folders/4_/33p_vn057msfh2nvgx6hwv_40000gn/T/$$chimera", nil
*/
func GetTempDir() (string, error) {
	if strKit.IsNotEmpty(tempDir) {
		return tempDir, nil
	}

	dir := Join(GetOsTempDir(), "$$"+consts.ProjectName)
	if err := fileKit.MkDirs(dir); err != nil {
		return "", err
	}
	return dir, nil
}

func SetTempDir(path string) error {
	if err := fileKit.AssertNotExistOrIsDir(path); err != nil {
		return err
	}

	tempDir = path
	return nil
}
