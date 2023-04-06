package ginKit

import (
	"github.com/gin-gonic/gin"
	"mime/multipart"
)

// UnsafeParamCapture 默认的，不安全的
var UnsafeParamCapture = &paramCapture{
	safe: false,
}

// SafeParamCapture 安全的，如果满足以下条件则必须使用此实例：请求转发、请求转发前会从POST中取参数、不会覆盖请求参数
var SafeParamCapture = &paramCapture{
	safe: true,
}

func GetParamFromGet(ctx *gin.Context, key string) string {
	return UnsafeParamCapture.GetParamFromGet(ctx, key)
}

func GetParamFromPost(ctx *gin.Context, key string) string {
	return UnsafeParamCapture.GetParamFromPost(ctx, key)
}

func GetParam(ctx *gin.Context, key string) string {
	return UnsafeParamCapture.GetParam(ctx, key)
}

func GetBoolParam(ctx *gin.Context, key string, def bool) bool {
	return UnsafeParamCapture.GetBoolParam(ctx, key, def)
}

func GetIntParam(ctx *gin.Context, key string, def int) int {
	return UnsafeParamCapture.GetIntParam(ctx, key, def)
}

//func GetInt32Param(ctx *gin.Context, key string, def int32) int32 {
//	return UnsafeParamCapture.GetInt32Param(ctx, key, def)
//}
//
//func GetInt64Param(ctx *gin.Context, key string, def int64) int64 {
//	return UnsafeParamCapture.GetInt64Param(ctx, key, def)
//}

func GetFloat32Param(ctx *gin.Context, key string, def float32) float32 {
	return UnsafeParamCapture.GetFloat32Param(ctx, key, def)
}

func GetFloat64Param(ctx *gin.Context, key string, def float64) float64 {
	return UnsafeParamCapture.GetFloat64Param(ctx, key, def)
}

func GetFormFile(ctx *gin.Context, key string) (*multipart.FileHeader, error) {
	return UnsafeParamCapture.GetFormFile(ctx, key)
}

func GetFormFileContent(ctx *gin.Context, key string) ([]byte, string, error) {
	return UnsafeParamCapture.GetFormFileContent(ctx, key)
}
