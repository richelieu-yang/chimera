package ginKit

import (
	"github.com/gin-gonic/gin"
	"github.com/richelieu-yang/chimera/v2/internal/resources"
	"github.com/richelieu-yang/chimera/v2/src/core/ioKit"
	"github.com/sirupsen/logrus"
	"net/http"
	"sync"
)

var noRouteData []byte
var noRouteErr error
var noRouteOnce sync.Once

func init() {
	var err error

	fs := resources.AssetFile()
	f, err := fs.Open("_resources/html/404.min.html")
	if err != nil {
		logrus.Fatal(err)
	}
	defer f.Close()

	noRouteData, err = ioKit.ReadFromReader(f)
	if err != nil {
		logrus.Fatal(err)
	}
}

// NoRoute 404
func NoRoute(engine IEngine, handlers ...gin.HandlerFunc) {
	engine.NoRoute(handlers...)
}

// DefaultNoRouteHtml 使用自带的404页面.
func DefaultNoRouteHtml(engine IEngine) error {
	noRouteOnce.Do(func() {
		path := "_resources/html/404.min.html"

		fs := resources.AssetFile()
		f, err := fs.Open(path)
		if err != nil {
			noRouteErr = err
			return
		}
		defer f.Close()

		noRouteData, err = ioKit.ReadFromReader(f)
		if err != nil {
			noRouteErr = err
			return
		}
	})

	if noRouteErr != nil {
		return noRouteErr
	}
	engine.NoRoute(func(ctx *gin.Context) {
		ctx.Data(http.StatusNotFound, "text/html; charset=utf-8", noRouteData)

		//// 此处不使用 ctx.FileFromFS ，原因: 这样的话，响应状态码就会是200，改不了
		//fs := resources.AssetFile()
		//ctx.FileFromFS("_resources/html/404.min.html", fs)
	})
	return nil
}
