package httpKit

import "net/http"

// GetContentType 获取 ContentType(即MimeType).
/*
@return 保底 "application/octet-stream"

e.g.
([]byte(nil))	=> "text/plain; charset=utf-8"
([]byte{}) 		=> "text/plain; charset=utf-8"
*/
func GetContentType(data []byte) string {
	return http.DetectContentType(data)
}
