package types

type JsonResponse struct {
	Code    string      `json:"code" example:"0"`
	Message string      `json:"message" example:"no error"`
	Data    interface{} `json:"data,omitempty"`
}
