package jsonKit

// Response
/*
响应给前端的json对象.
*/
type Response struct {
	Code    string      `json:"code" example:"0"`
	Message string      `json:"message" example:"no error"`
	Data    interface{} `json:"data,omitempty"`
}
