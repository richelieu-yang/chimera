package httpClientKit

import (
	"net/http"
)

// send 通用代码.
/*
@return 第一个返回值如果不为nil的话，一般来说需要手动调用 "resp.Body.Close()" !!!
*/
func send(client *http.Client, req *http.Request) (*http.Response, error) {
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
