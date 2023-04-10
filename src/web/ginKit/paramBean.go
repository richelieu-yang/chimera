package ginKit

import (
	"github.com/gin-gonic/gin"
	"github.com/richelieu42/chimera/v2/src/core/boolKit"
	"github.com/richelieu42/chimera/v2/src/core/floatKit"
	"github.com/richelieu42/chimera/v2/src/core/intKit"
	"github.com/richelieu42/chimera/v2/src/web/httpKit"
	"io"
	"mime/multipart"
	"strings"
)

type paramCapture struct {
	safe bool
}

// GetParam 从请求中获取参数（优先通过GET请求方式）
func (c *paramCapture) GetParam(ctx *gin.Context, key string) string {
	// from GET
	str := c.GetParamFromGet(ctx, key)
	if str != "" {
		return str
	}
	// from POST
	return c.GetParamFromPost(ctx, key)
}

// GetParamFromGet 从GET请求中获取参数.
/*
3种方法:
	str = c.Query("wd")
	str := c.DefaultQuery("wd","acwing")
	str , ok := c.GetQuery("wd")
*/
func (c *paramCapture) GetParamFromGet(ctx *gin.Context, key string) string {
	return ctx.Query(key)
}

// GetParamFromPost 从 "普通POST请求"(x-www-form-urlencoded) 、 "form表单提交"(form-data) 中获取参数.
/*
3种方法:
	username := c.PostForm("username")
	username = c.DefaultPostForm("username","somebody")
	username , ok := c.GetPostForm("username")
*/
func (c *paramCapture) GetParamFromPost(ctx *gin.Context, key string) string {
	if c.safe {
		// Golang中比较坑的一点：如果不通过此方法获取post参数，如果转发请求，接收方取参数会有问题.
		req := ctx.Request

		// 重构body
		var bodyBytes []byte
		if req.Body != nil {
			bodyBytes, _ = io.ReadAll(req.Body)
		}
		reader := strings.NewReader(string(bodyBytes))
		req.Body = &httpKit.Repeat{Reader: reader, Offset: 0}

		str := ctx.PostForm(key)

		// 重置body（否则转发请求会失败）
		req.Body.(*httpKit.Repeat).Reset()

		return str
	}
	return ctx.PostForm(key)
}

func (c *paramCapture) GetBoolParam(ctx *gin.Context, key string, def bool) bool {
	str := c.GetParam(ctx, key)
	return boolKit.ToBoolWithDefault(str, def)
}

func (c *paramCapture) GetIntParam(ctx *gin.Context, key string, def int) int {
	str := c.GetParam(ctx, key)
	return intKit.StringToIntWithDefault(str, def)
}

func (c *paramCapture) GetFloat32Param(ctx *gin.Context, key string, def float32) float32 {
	str := c.GetParam(ctx, key)
	return floatKit.ToFloat32WithDefault(str, def)
}

func (c *paramCapture) GetFloat64Param(ctx *gin.Context, key string, def float64) float64 {
	str := c.GetParam(ctx, key)
	return floatKit.ToFloat64WithDefault(str, def)
}

// GetFormFile form请求，根据 传参key 获取文件
/*
@param key 	可以为""
@return 	如果不存在与 key 对应的文件，将返回error（http: no such file）
*/
func (c *paramCapture) GetFormFile(ctx *gin.Context, key string) (*multipart.FileHeader, error) {
	return ctx.FormFile(key)
}

// GetFormFileContent form请求，根据 传参key 获取文件的字节流
/*
@return 分别为：文件内容、文件名、错误
*/
func (c *paramCapture) GetFormFileContent(ctx *gin.Context, key string) ([]byte, string, error) {
	fileHeader, err := c.GetFormFile(ctx, key)
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
