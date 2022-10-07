package httpClientKit

import (
	"gitee.com/richelieu042/go-scales/src/core/errorKit"
	"net/http"
)

// AssertHttpStatusCodeSuccessful 断言http状态码是否成功
/*
PS: 判断参考了jQuery.
*/
func AssertHttpStatusCodeSuccessful(statusCode int) error {
	if statusCode >= 200 && statusCode < 300 || statusCode == http.StatusNotModified {
		return nil
	}
	// 此处的1是为了跳过当前函数的调用
	return errorKit.SimpleWithExtraSkip(1, "[Assertion failed] http status code(%d) isn't successful", statusCode)
}
