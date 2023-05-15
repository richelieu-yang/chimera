package mimeTypeKit

import (
	"github.com/gabriel-vasile/mimetype"
	"io"
	"net/http"
)

// DetectContentType 获取 ContentType(即MimeType).
/*
@return 保底 "application/octet-stream"

e.g.
([]byte(nil))	=> "text/plain; charset=utf-8"
([]byte{}) 		=> "text/plain; charset=utf-8"
*/
func DetectContentType(data []byte) string {
	return http.DetectContentType(data)
}

// Detect
/*
mimetype库: 基于magic数的用于媒体类型和文件扩展名检测的快速的 Go 库，支持 170+ 格式.
*/
func Detect(in []byte) *mimetype.MIME {
	return mimetype.Detect(in)
}

func DetectReader(r io.Reader) (*mimetype.MIME, error) {
	return mimetype.DetectReader(r)
}

func DetectFile(path string) (*mimetype.MIME, error) {
	return mimetype.DetectFile(path)
}
