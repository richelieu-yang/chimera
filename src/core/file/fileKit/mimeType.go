package fileKit

import "net/http"

// GetMimeType 获取 MimeType
/*
@return 保底 "application/octet-stream"
*/
func GetMimeType(data []byte) string {
	return http.DetectContentType(data)
}
