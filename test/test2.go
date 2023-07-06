package main

import (
	"github.com/gin-gonic/gin"
	"github.com/richelieu-yang/chimera/v2/src/log/logrusKit"
	"github.com/richelieu-yang/chimera/v2/src/web/cookieKit"
	"github.com/sirupsen/logrus"
	"net/http"
)

func main() {
	logrusKit.MustSetUp(&logrusKit.Config{
		Level:      "",
		PrintBasic: false,
	})
	engine := gin.Default()

	engine.Any("/test", func(ctx *gin.Context) {
		cookie := cookieKit.NewCookie("cyy", "yjs", "", "", 0, false, true, http.SameSiteDefaultMode)
		cookieKit.SetCookie(ctx.Writer, cookie)
		ctx.String(http.StatusOK, "ok")
	})
	engine.Static("/s/", "./resource2")

	if err := engine.Run(":82"); err != nil {
		logrus.Fatal(err)
	}
}
