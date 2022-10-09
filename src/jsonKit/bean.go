package jsonKit

// JsonResponse
/*
响应给前端的json对象.
*/
type JsonResponse struct {
	Code    string      `json:"code" example:"0"`
	Message string      `json:"message" example:"no error"`
	Data    interface{} `json:"data,omitempty"`
}
