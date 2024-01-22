package ginKit

import (
	"github.com/gin-gonic/gin"
	"github.com/richelieu-yang/chimera/v2/src/component/web/httpKit"
)

// SetHttpStatusCode （响应头）设置http状态码
func SetHttpStatusCode(ctx *gin.Context, statusCode int) {
	ctx.Writer.WriteHeader(statusCode)
}

// GetRequestHeader （请求头）
func GetRequestHeader(ctx *gin.Context, key string) string {
	return ctx.GetHeader(key)
}

// SetResponseHeader （响应头）
func SetResponseHeader(ctx *gin.Context, key, value string) {
	httpKit.SetHeader(ctx.Writer.Header(), key, value)
}

// AddResponseHeader （响应头）
func AddResponseHeader(ctx *gin.Context, key, value string) {
	httpKit.AddHeader(ctx.Writer.Header(), key, value)
}

// DelResponseHeader （响应头）删除响应头.
func DelResponseHeader(ctx *gin.Context, key string) {
	httpKit.DelHeader(ctx.Writer.Header(), key)
}

// SetCacheControlNoCache 实际上是有缓存的）浏览器对请求回来的response做缓存，但是每次在向客户端（浏览器）提供响应数据时，缓存都要向服务器评估缓存响应的有效性。
/*
PS: 详见"Web.docx".
*/
func SetCacheControlNoCache(ctx *gin.Context) {
	SetResponseHeader(ctx, "Cache-Control", "no-cache")
}

// SetCacheControlNoStore 禁止一切缓存.
/*
PS: 详见"Web.docx".
*/
func SetCacheControlNoStore(ctx *gin.Context) {
	SetResponseHeader(ctx, "Cache-Control", "no-store")
}

// GetUserAgent 获取http请求头中"User Agent"的值.
/*
参考: https://www.sunzhongwei.com/golang-gin-for-user-agent-in-http-request-header-value

e.g.
Chrome浏览器: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/103.0.0.0 Safari/537.36
Safari浏览器: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/15.5 Safari/605.1.15
*/
func GetUserAgent(ctx *gin.Context) string {
	return GetRequestHeader(ctx, "User-Agent")
}
