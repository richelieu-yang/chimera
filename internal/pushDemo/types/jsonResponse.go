package types

type JsonResponse struct {
	Code    string      `json:"code" example:"0"`
	Message string      `json:"message" example:"No error"`
	Data    interface{} `json:"data,omitempty"`
}
