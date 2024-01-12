package compressKit

import (
	"github.com/gogf/gf/v2/encoding/gcompress"
)

func UnGzip(data []byte) ([]byte, error) {
	return gcompress.UnGzip(data)
}

func UnGzipFile(srcFilePath, dstFilePath string) error {
	return gcompress.UnGzipFile(srcFilePath, dstFilePath)
}
