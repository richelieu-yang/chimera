package ginKit

import (
	"github.com/gin-gonic/gin"
	"github.com/richelieu42/chimera/v2/src/core/boolKit"
	"github.com/richelieu42/chimera/v2/src/core/floatKit"
	"github.com/richelieu42/chimera/v2/src/core/intKit"
	"github.com/richelieu42/chimera/v2/src/core/sliceKit"
	"github.com/richelieu42/chimera/v2/src/core/strKit"
	"github.com/richelieu42/chimera/v2/src/web/httpKit"
	"io"
	"strings"
)

func ObtainGetParam(ctx *gin.Context, key string) string {
	return ctx.Query(key)
}

func ObtainPostParam(ctx *gin.Context, key string, safeArgs ...bool) string {
	safe := sliceKit.GetFirstItemWithDefault(false, safeArgs...)
	if safe && ctx.Request.Body != nil {
		/* Golang中比较坑的一点：如果不通过此方法获取post参数，如果涉及转发请求，接收方取参数会有问题. */
		// (1) 重构body
		bodyBytes, _ := io.ReadAll(ctx.Request.Body)
		//if err != nil {
		//	return nil, "", err
		//}
		reader := strings.NewReader(string(bodyBytes))
		ctx.Request.Body = &httpKit.Repeat{
			Reader: reader,
			Offset: 0,
		}
		defer func() {
			// (2) 重置body（否则转发请求会失败）
			ctx.Request.Body.(*httpKit.Repeat).Reset()
		}()
	}

	return ctx.PostForm(key)
}

// ObtainParam 获取请求参数（优先级: GET > POST）
func ObtainParam(ctx *gin.Context, key string, safeArgs ...bool) string {
	// (1) GET
	value := ObtainGetParam(ctx, key)
	if strKit.IsNotEmpty(value) {
		return value
	}

	// (2) POST
	return ObtainPostParam(ctx, key, safeArgs...)
}

func ObtainBoolParam(ctx *gin.Context, key string, safeArgs ...bool) (bool, error) {
	value := ObtainParam(ctx, key, safeArgs...)
	return boolKit.ToBoolE(value)
}

func ObtainIntParam(ctx *gin.Context, key string, safeArgs ...bool) (int, error) {
	value := ObtainParam(ctx, key, safeArgs...)
	return intKit.ToIntE(value)
}

func ObtainInt32Param(ctx *gin.Context, key string, safeArgs ...bool) (int32, error) {
	value := ObtainParam(ctx, key, safeArgs...)
	return intKit.ToInt32E(value)
}

func ObtainInt64Param(ctx *gin.Context, key string, safeArgs ...bool) (int64, error) {
	value := ObtainParam(ctx, key, safeArgs...)
	return intKit.ToInt64E(value)
}

func ObtainFloat32Param(ctx *gin.Context, key string, safeArgs ...bool) (float32, error) {
	value := ObtainParam(ctx, key, safeArgs...)
	return floatKit.ToFloat32E(value)
}

func ObtainFloat64Param(ctx *gin.Context, key string, safeArgs ...bool) (float64, error) {
	value := ObtainParam(ctx, key, safeArgs...)
	return floatKit.ToFloat64E(value)
}

// ObtainFormFileContent form请求，根据 传参key 获取文件的字节流
/*
TODO: 待验证：form表单上传文件，传参safeArgs是否有必要?

@return 文件内容 + 文件名 + 错误
*/
func ObtainFormFileContent(ctx *gin.Context, key string, safeArgs ...bool) ([]byte, string, error) {
	safe := sliceKit.GetFirstItemWithDefault(false, safeArgs...)
	if safe && ctx.Request.Body != nil {
		/* Golang中比较坑的一点：如果不通过此方法获取post参数，如果涉及转发请求，接收方取参数会有问题. */
		// (1) 重构body
		bodyBytes, err := io.ReadAll(ctx.Request.Body)
		if err != nil {
			return nil, "", err
		}
		reader := strings.NewReader(string(bodyBytes))
		ctx.Request.Body = &httpKit.Repeat{
			Reader: reader,
			Offset: 0,
		}
		defer func() {
			// (2) 重置body（否则转发请求会失败）
			ctx.Request.Body.(*httpKit.Repeat).Reset()
		}()
	}

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
