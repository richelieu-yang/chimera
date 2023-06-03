package pathKit

import (
	"github.com/richelieu-yang/chimera/v2/src/consts"
	"github.com/richelieu-yang/chimera/v2/src/core/fileKit"
	"os"
)

// GetTempDir 获取系统的临时目录.
/*
e.g.
Windows: 	"C:\Users\Lenovo\AppData\Local\Temp"
Mac: 		"/var/folders/4_/33p_vn057msfh2nvgx6hwv_40000gn/T/"
*/
func GetTempDir() string {
	return os.TempDir()
}

// GetUniqueTempDir 获取 本依赖 的专属临时目录.
/*
e.g. Mac
() => "/var/folders/4_/33p_vn057msfh2nvgx6hwv_40000gn/T/$$chimera", nil
*/
func GetUniqueTempDir() (string, error) {
	dir := Join(GetTempDir(), "$$"+consts.ProjectName)
	err := fileKit.MkDirs(dir)
	return dir, err
}
