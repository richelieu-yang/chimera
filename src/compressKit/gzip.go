package compressKit

import (
	"github.com/gogf/gf/v2/encoding/gcompress"
	"io"
)

func Gzip(data []byte, level ...int) ([]byte, error) {
	return gcompress.Gzip(data, level...)
}

func GzipFile(srcFilePath, dstFilePath string, level ...int) (err error) {
	return gcompress.GzipFile(srcFilePath, dstFilePath, level...)
}

func GzipPathWriter(filePath string, writer io.Writer, level ...int) error {
	return gcompress.GzipPathWriter(filePath, writer, level...)
}

func UnGzip(data []byte) ([]byte, error) {
	return gcompress.UnGzip(data)
}

func UnGzipFile(srcFilePath, dstFilePath string) error {
	return gcompress.UnGzipFile(srcFilePath, dstFilePath)
}
