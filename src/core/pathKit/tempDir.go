package pathKit

import (
	"github.com/richelieu42/go-scales/src/core/file/fileKit"
	"os"
)

// GetTempDir 默认返回: 系统的临时目录
/*
e.g.
Windows: 	"C:\Users\Lenovo\AppData\Local\Temp"
Mac: 		"/var/folders/4_/33p_vn057msfh2nvgx6hwv_40000gn/T/"
*/
func GetTempDir() string {
	return os.TempDir()
}

// GetTempDirOfGoScales go-scales专属的临时目录
func GetTempDirOfGoScales() (string, error) {
	dir := Join(GetTempDir(), "$go-scales")
	err := fileKit.MkDirs(dir)
	return dir, err
}
