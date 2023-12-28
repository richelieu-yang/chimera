package zipKit

import "github.com/duke-git/lancet/v2/fileutil"

var (
	Zip func(path string, destPath string) error = fileutil.Zip

	UnZip func(zipFile string, destPath string) error = fileutil.UnZip

	ZipAppendEntry func(fpath string, destPath string) error = fileutil.ZipAppendEntry

	IsZipFile func(filepath string) bool = fileutil.IsZipFile
)
