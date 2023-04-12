package httpKit

import "net/http"

// LimitRequestBody 限制请求body的大小
/*
PS: 适用于限制上传文件的大小
*/
func LimitRequestBody(w http.ResponseWriter, r *http.Request, n int64) {
	r.Body = http.MaxBytesReader(w, r.Body, n)
}
