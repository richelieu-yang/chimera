package httpClientKit

import (
	"bytes"
	"github.com/richelieu42/go-scales/src/core/file/fileKit"
	"github.com/richelieu42/go-scales/src/core/mapKit"
	"github.com/richelieu42/go-scales/src/core/strKit"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"strings"
)

// SimplePost 额外处理了http状态码
func SimplePost(url string, params map[string]string) ([]byte, error) {
	statusCode, data, err := Post(url, params)
	if err != nil {
		return nil, err
	}
	if err = AssertHttpStatusCodeSuccessful(statusCode); err != nil {
		return nil, err
	}
	return data, nil
}

// Post
/*
@param params 	请求参数，可以为nil
@return 		http状态码 + 响应内容 + error
*/
func Post(url string, params map[string]string) (int, []byte, error) {
	// body
	paramStr := mapKit.JoinSS(params, "&", "=", nil, nil)
	body := strings.NewReader(paramStr)

	return post(url, body, "application/x-www-form-urlencoded;charset=utf-8")
}

// UploadFile 上传单个文件
/*
TODO: 待验证及测试，以及和上面的方法一起抽出通用代码.

@param params 可以为nil
*/
func UploadFile(url string, params map[string]string, fileKey, filePath string) (statusCode int, data []byte, err error) {
	if err = fileKit.AssertExistAndIsFile(filePath); err != nil {
		return
	}

	// fileKey默认值: "file"
	fileKey = strKit.EmptyToDefault(fileKey, "file")

	// body
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	for k, v := range params {
		// 此处无需对k、v进行编码处理
		err = writer.WriteField(k, v)
		if err != nil {
			return
		}
	}
	file, err := os.Open(filePath)
	if err != nil {
		return
	}
	defer file.Close()
	formPart, err := writer.CreateFormFile(fileKey, fileKit.GetName(filePath))
	if err != nil {
		return
	}
	_, err = io.Copy(formPart, file)
	if err != nil {
		return
	}
	err = writer.Close()
	if err != nil {
		return
	}

	return post(url, body, writer.FormDataContentType())
}

// post
/*
@param body			可以为nil
@param contentType	可以为""
*/
func post(url string, body io.Reader, contentType string) (statusCode int, data []byte, err error) {
	if err = strKit.AssertNotBlank(url, "url"); err != nil {
		return
	}

	client := newClient()
	req, err := http.NewRequest("POST", url, body)
	if err != nil {
		return
	}
	if strKit.IsNotEmpty(contentType) {
		req.Header.Add("Content-Type", contentType)
	}
	resp, err := client.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	data, err = io.ReadAll(resp.Body)
	statusCode = resp.StatusCode
	return
}
