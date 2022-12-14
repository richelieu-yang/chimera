package httpClientKit

import (
	"github.com/richelieu42/go-scales/src/core/mapKit"
	"github.com/richelieu42/go-scales/src/core/strKit"
	"github.com/richelieu42/go-scales/src/urlKit"
	"io"
	"net/http"
)

// SimpleGet 额外处理了http状态码
func SimpleGet(url string, params map[string]string) ([]byte, error) {
	statusCode, data, err := Get(url, params)
	if err != nil {
		return nil, err
	}
	if err = AssertHttpStatusCodeSuccessful(statusCode); err != nil {
		return nil, err
	}
	return data, nil
}

// Get 发送GET请求（可用于下载文件等场景）
/*
@param params 	请求参数，可以为nil
@return 		http状态码 + 响应内容 + error

参考: golang 将图片生成Base64 https://blog.csdn.net/weixin_40292098/article/details/126029489
*/
func Get(url string, params map[string]string) (int, []byte, error) {
	// 参数加到url上
	callback := func(str string) string {
		return urlKit.EncodeURIComponent(str)
	}
	paramStr := mapKit.JoinSS(params, "&", "=", callback, callback)
	if strKit.IsNotEmpty(paramStr) {
		count := strKit.Count(url, "?")
		switch count {
		case 0:
			url += "?" + paramStr
		case 1:
			if strKit.EndWith(url, "?") {
				url += paramStr
			} else {
				url += "&" + paramStr
			}
		default:
			url += "&" + paramStr
		}
	}

	// 发请求
	return get(url)
}

func get(url string) (statusCode int, data []byte, err error) {
	if err = strKit.AssertNotBlank(url, "url"); err != nil {
		return
	}

	client := newClient()
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return
	}
	resp, err := client.Do(request)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	data, err = io.ReadAll(resp.Body)
	if err != nil {
		return
	}
	statusCode = resp.StatusCode
	return
}
