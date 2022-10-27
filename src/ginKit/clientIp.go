package ginKit

import "github.com/gin-gonic/gin"

// GetClientIp 获取客户端的ip
/*
参考：
gin框架中设置信任代理IP并获取远程客户端IP
	https://www.cnblogs.com/mayanan/p/15703234.html
gin获取用户请求IP
	https://blog.csdn.net/weixin_45867397/article/details/122849424

PS: Context.ClientIP() 可以和 Engine.TrustedPlatform 搭配使用，详见参考.
*/
func GetClientIp(ctx *gin.Context) string {
	/*
		Context.RemoteIP(): 无代理返回客户端IP,有代理返回代理IP
		Context.ClientIP(): 无论是否有代理，都会返回客户端IP（代理比如Nginx也需要配置；方法体内部会调用 Context.RemoteIP()）
	*/
	return ctx.ClientIP()
}
