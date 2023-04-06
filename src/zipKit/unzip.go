package zipKit

import (
	"github.com/dablelv/go-huge-util/zip"
)

// Unzip 解压
/*
@param zipPath e.g."archive.zip"
*/
func Unzip(zipPath, dstDir string) error {
	return zip.Unzip(zipPath, dstDir)
}
