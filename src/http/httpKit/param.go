package httpKit

import (
	"mime/multipart"
	"net/http"
)

var (
	// defaultMaxMemory 默认文件大小上限: 10 MB
	defaultMaxMemory int64 = 10 << 20
)

// GetUrlParam
/*
[Go] golang获取http中的get传递的参数
	https://www.cnblogs.com/taoshihan/p/12943118.html
*/
func GetUrlParam(r *http.Request, key string) string {
	values := r.URL.Query()
	return values.Get(key)
}

// GetUrlParam1
/*
[Go] golang获取http中的get传递的参数
	https://www.cnblogs.com/taoshihan/p/12943118.html
*/
func GetUrlParam1(r *http.Request, key string) (string, error) {
	if err := r.ParseForm(); err != nil {
		return "", err
	}
	return r.FormValue(key), nil
}

// GetPostParam
/*
[Go] golang获取http中的get传递的参数
	https://www.cnblogs.com/taoshihan/p/12943118.html
*/
func GetPostParam(r *http.Request, key string) (string, error) {
	// 解析: application/x-www-form-urlencoded
	if err := r.ParseForm(); err != nil {
		return "", err
	}
	return r.PostFormValue(key), nil
}

// GetFormFile
/*
[Go] golang获取http中的get传递的参数
	https://www.cnblogs.com/taoshihan/p/12943118.html

@param maxMemory 可以为-1，此时将采用默认值
*/
func GetFormFile(r *http.Request, key string, maxMemory int64) (multipart.File, *multipart.FileHeader, error) {
	if maxMemory < 0 {
		maxMemory = defaultMaxMemory
	}

	// 解析: multipart/form-data
	if err := r.ParseMultipartForm(maxMemory); err != nil {
		return nil, nil, err
	}
	return r.FormFile(key)
}
