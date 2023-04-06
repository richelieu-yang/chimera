// Package httpStatusCode 自定义的http状态码
/*
ajax success支持的状态码（https://www.cnblogs.com/52liming/p/9537289.html）：isSuccess = status >= 200 && status < 300 || status === 304;

HTTP状态码总的分为5类（https://www.runoob.com/http/http-status-codes.html）：
1开头：信息状态码
2开头：成功状态码
3开头：重定向状态码
4开头：客户端错误状态码
5开头：服务端错误状态码
*/
package httpStatusCode

import "net/http"

const (
	OK = http.StatusOK /*200*/

	NotModified = http.StatusNotModified /*304*/
	// FailToForward 请求转发失败
	FailToForward = 380

	NotFound = http.StatusNotFound /*404*/
	// ServiceNotFound 服务不存在（比如：请求转发的目标服务不存在）
	ServiceNotFound = 480
	// ResourceNotFound 资源不存在（为了和默认的404做区分）
	ResourceNotFound = 490
	// ResourceIsRemoved 资源已经被移除
	ResourceIsRemoved = 491

	// Panic 服务器端发生panic
	Panic = 520
)
