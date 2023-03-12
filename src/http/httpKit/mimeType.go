package httpKit

import "net/http"

// GetMimeType 获取 MimeType
/*
@return 保底 "application/octet-stream"

e.g.
([]byte{}) => "text/plain; charset=utf-8"
*/
func GetMimeType(data []byte) string {
	return http.DetectContentType(data)
}
