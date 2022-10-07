package zipKit

import (
	huge "github.com/dablelv/go-huge-util"
)

// Unzip 解压
/*
@param zipPath e.g."archive.zip"
*/
func Unzip(zipPath, dstDir string) error {
	return huge.Unzip(zipPath, dstDir)
}
