package httpKit

import "net/http"

// SetCacheControlNoCache 实际上是有缓存的）浏览器对请求回来的response做缓存，但是每次在向客户端（浏览器）提供响应数据时，缓存都要向服务器评估缓存响应的有效性。
/*
PS:
(1) 一般情况下， "no-cache" 就够了；
(2) 详见"Web.docx".
*/
func SetCacheControlNoCache(header http.Header) {
	SetHeader(header, "Cache-Control", "no-cache")
}

// SetCacheControlNoStore 禁止一切缓存.
/*
PS: 详见"Web.docx".
*/
func SetCacheControlNoStore(header http.Header) {
	SetHeader(header, "Cache-Control", "no-store")
}
