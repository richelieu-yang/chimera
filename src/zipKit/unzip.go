package zipKit

import "github.com/dablelv/cyan/zip"

// Unzip 解压
/*
@param zipPath e.g."archive.zip"
*/
var Unzip func(zipPath, dstDir string) error = zip.Unzip
