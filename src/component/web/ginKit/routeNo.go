package ginKit

import (
	"github.com/gin-gonic/gin"
	"github.com/richelieu-yang/chimera/v2/internal/resources"
	"github.com/richelieu-yang/chimera/v2/src/component/web/httpKit"
	"github.com/richelieu-yang/chimera/v2/src/core/fileKit"
	"github.com/richelieu-yang/chimera/v2/src/core/pathKit"
	"net/http"
)

// NoRoute 404
func NoRoute(engine IEngine, handlers ...gin.HandlerFunc) {
	engine.NoRoute(handlers...)
}

// AttachDefaultNoRoute 使用自带的404页面.
/*
PS: 会在临时目录下生成文件，注意不要删掉他们!!!
*/
func AttachDefaultNoRoute(engine IEngine) error {
	/* 将内置的404页面解压到 临时目录 中 */
	relativePath := "_resources/html/404.html"
	tempDir, err := pathKit.GetTempDir()
	if err != nil {
		return err
	}
	err = resources.RestoreAsset(tempDir, relativePath)
	if err != nil {
		return err
	}

	/* 加载解压出的html页面 */
	htmlPath := pathKit.Join(tempDir, relativePath)
	engine.LoadHTMLFiles(htmlPath)

	NoRoute(engine, func(ctx *gin.Context) {
		ctx.HTML(http.StatusNotFound, fileKit.GetFileName(htmlPath), gin.H{
			"route": httpKit.GetRoute(ctx.Request),
		})
	})

	return nil
}
