package main

import (
	"github.com/gin-gonic/gin"
	"github.com/richelieu-yang/chimera/v2/src/component/web/ginKit"
)

func main() {
	engine := gin.Default()

	if err := ginKit.DefaultNoRoute(engine); err != nil {
		panic(err)
	}

	//fs := resources.AssetFile()
	////engine.StaticFileFS("/404", "_resources/html/404.min.html", fs)
	//engine.NoRoute(func(ctx *gin.Context) {
	//	f, err := fs.Open("_resources/html/404.min.html")
	//	if err != nil {
	//		panic(err)
	//	}
	//	defer f.Close()
	//
	//	data, err := ioKit.ReadFromReader(f)
	//	if err != nil {
	//		panic(err)
	//	}
	//	ctx.Data(http.StatusNotFound, "text/html; charset=utf-8", data)
	//
	//	//ctx.FileFromFS("_resources/html/404.min.html", fs)
	//	//ctx.Status(404)
	//})

	if err := engine.Run(":80"); err != nil {
		panic(err)
	}
}
