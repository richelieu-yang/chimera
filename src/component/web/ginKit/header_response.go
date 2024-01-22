package ginKit

import (
	"github.com/gin-gonic/gin"
	"github.com/richelieu-yang/chimera/v2/src/component/web/httpKit"
)

// SetHttpStatusCode （响应头）设置http状态码
func SetHttpStatusCode(ctx *gin.Context, statusCode int) {
	ctx.Writer.WriteHeader(statusCode)
}

// SetResponseHeader （响应头）
func SetResponseHeader(ctx *gin.Context, key, value string) {
	ctx.Header(key, value)
	//httpKit.SetHeader(ctx.Writer.Header(), key, value)
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
