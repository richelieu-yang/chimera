package compressKit

import (
	"github.com/gogf/gf/v2/encoding/gcompress"
	"io"
)

func ZipPath(fileOrFolderPaths, dstFilePath string, prefix ...string) error {
	return gcompress.ZipPath(fileOrFolderPaths, dstFilePath, prefix...)
}

func ZipPathContent(fileOrFolderPaths string, prefix ...string) ([]byte, error) {
	return gcompress.ZipPathContent(fileOrFolderPaths, prefix...)
}

func ZipPathWriter(fileOrFolderPaths string, writer io.Writer, prefix ...string) error {
	return gcompress.ZipPathWriter(fileOrFolderPaths, writer, prefix...)
}

func UnZipContent(zippedContent []byte, dstFolderPath string, zippedPrefix ...string) error {
	return gcompress.UnZipContent(zippedContent, dstFolderPath, zippedPrefix...)
}

func UnZipFile(zippedFilePath, dstFolderPath string, zippedPrefix ...string) error {
	return gcompress.UnZipFile(zippedFilePath, dstFolderPath, zippedPrefix...)
}
