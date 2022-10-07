// Package ginKit
/*
参考: 使用gin搭建api后台系统之cookie与session https://mp.weixin.qq.com/s/nSu4ZQ9lStvGFNnuC2WKnA
*/
package ginKit

import "github.com/gin-gonic/gin"

// GetCookie 读取cookie
/*
PS: 如果不存在与 name 对应的cookie，将返回error(== http.ErrNoCookie).
*/
func GetCookie(ctx *gin.Context, name string) (string, error) {
	return ctx.Cookie(name)
}

// SetCookie cookie
/*
@param name		cookie的键
@param value	cookie的值
@param maxAge	cookie的有效时长（单位为秒；超时即会失效）
@param path		cookie所在的目录（可以为""，会被"/"替换）
@param domain	所在域，表示我们的 cookie 作用范围，里面可以是localhost也可以是你的域名，看自己情况（可以为""，此时响应头中Set-Cookie将不会带有domain信息，前端自行生成与当前页面对应的）
@param secure	是否只能通过https访问？true: 只能通过https访问
@param httpOnly	cookie是否可以通过js代码进行操作？true: 不能被js进行操作
*/
func SetCookie(ctx *gin.Context, name, value string, maxAge int, path, domain string, secure, httpOnly bool) {
	ctx.SetCookie(name, value, maxAge, path, domain, secure, httpOnly)
}

// DelCookie 删除cookie
func DelCookie() {

}
