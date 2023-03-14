package httpKit

import (
	"mime/multipart"
	"net/http"
)

// GetUrlParam
/*
[Go] golang获取http中的get传递的参数
	https://www.cnblogs.com/taoshihan/p/12943118.html
*/
func GetUrlParam(r *http.Request, name string) string {
	values := r.URL.Query()
	return values.Get(name)
}

// GetUrlParam1
/*
[Go] golang获取http中的get传递的参数
	https://www.cnblogs.com/taoshihan/p/12943118.html
*/
func GetUrlParam1(r *http.Request, name string) (string, error) {
	if err := r.ParseForm(); err != nil {
		return "", err
	}
	return r.FormValue(name), nil
}

// GetPostParam
/*
[Go] golang获取http中的get传递的参数
	https://www.cnblogs.com/taoshihan/p/12943118.html
*/
func GetPostParam(r *http.Request, name string) (string, error) {
	if err := r.ParseForm(); err != nil {
		return "", err
	}
	return r.PostFormValue(name), nil
}

// GetFormFile
/*
[Go] golang获取http中的get传递的参数
	https://www.cnblogs.com/taoshihan/p/12943118.html
*/
func GetFormFile(r *http.Request, name string) (multipart.File, *multipart.FileHeader, error) {
	if err := r.ParseForm(); err != nil {
		return nil, nil, err
	}
	return r.FormFile(name)
}
