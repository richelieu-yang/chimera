package jsonKit

// Response
/*
响应给前端的json对象.
*/
type Response struct {
	Code    string      `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}
