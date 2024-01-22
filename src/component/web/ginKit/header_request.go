package ginKit

import "github.com/gin-gonic/gin"

// GetRequestHeader （请求头）
func GetRequestHeader(ctx *gin.Context, key string) string {
	return ctx.GetHeader(key)
}

// GetUserAgent （请求头）获取http请求头中"User Agent"的值.
/*
参考: https://www.sunzhongwei.com/golang-gin-for-user-agent-in-http-request-header-value

e.g.
Chrome浏览器: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/103.0.0.0 Safari/537.36
Safari浏览器: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/15.5 Safari/605.1.15
*/
func GetUserAgent(ctx *gin.Context) string {
	return GetRequestHeader(ctx, "User-Agent")
}
