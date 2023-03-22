package pathKit

import (
	"github.com/richelieu42/chimera/src/core/file/fileKit"
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

// GetChimeraTempDir 获取chimera的专属临时目录
/*
e.g. Mac
() => "/var/folders/4_/33p_vn057msfh2nvgx6hwv_40000gn/T/$$chimera", nil
*/
func GetChimeraTempDir() (string, error) {
	dir := Join(GetTempDir(), "$$chimera")
	err := fileKit.MkDirs(dir)
	return dir, err
}
