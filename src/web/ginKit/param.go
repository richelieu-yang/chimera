package ginKit

import (
	"github.com/gin-gonic/gin"
	"github.com/richelieu42/chimera/v2/src/core/boolKit"
	"github.com/richelieu42/chimera/v2/src/core/floatKit"
	"github.com/richelieu42/chimera/v2/src/core/intKit"
	"github.com/richelieu42/chimera/v2/src/core/strKit"
	"io"
)

func ObtainGetParam(ctx *gin.Context, key string) string {
	return ctx.Query(key)
}

func ObtainPostParam(ctx *gin.Context, key string) string {
	return ctx.PostForm(key)
}

// ObtainParam 获取请求参数（优先级: GET > POST）
func ObtainParam(ctx *gin.Context, key string) string {
	// (1) GET
	value := ObtainGetParam(ctx, key)
	if strKit.IsNotEmpty(value) {
		return value
	}

	// (2) POST
	return ObtainPostParam(ctx, key)
}

func ObtainBoolParam(ctx *gin.Context, key string) (bool, error) {
	value := ObtainParam(ctx, key)
	return boolKit.ToBoolE(value)
}

func ObtainIntParam(ctx *gin.Context, key string) (int, error) {
	value := ObtainParam(ctx, key)
	return intKit.ToIntE(value)
}

func ObtainInt32Param(ctx *gin.Context, key string) (int32, error) {
	value := ObtainParam(ctx, key)
	return intKit.ToInt32E(value)
}

func ObtainInt64Param(ctx *gin.Context, key string) (int64, error) {
	value := ObtainParam(ctx, key)
	return intKit.ToInt64E(value)
}

func ObtainFloat32Param(ctx *gin.Context, key string) (float32, error) {
	value := ObtainParam(ctx, key)
	return floatKit.ToFloat32E(value)
}

func ObtainFloat64Param(ctx *gin.Context, key string) (float64, error) {
	value := ObtainParam(ctx, key)
	return floatKit.ToFloat64E(value)
}

// ObtainFormFileContent form请求，根据 传参key 获取文件的字节流
/*
@return 文件内容 + 文件名 + 错误
*/
func ObtainFormFileContent(ctx *gin.Context, key string) ([]byte, string, error) {
	fileHeader, err := ctx.FormFile(key)
	if err != nil {
		return nil, "", err
	}

	file, err := fileHeader.Open()
	if err != nil {
		return nil, "", err
	}
	defer file.Close()
	content, err := io.ReadAll(file)
	if err != nil {
		return nil, "", err
	}
	return content, fileHeader.Filename, nil
}
