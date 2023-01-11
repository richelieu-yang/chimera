package main

import (
	"github.com/gin-gonic/gin"
	"github.com/richelieu42/go-scales/src/cookieKit"
	"github.com/richelieu42/go-scales/src/urlKit"
	"net/http"
	"strconv"
)

var name = "name"

func main() {
	engine := gin.Default()

	engine.Any("/set", func(ctx *gin.Context) {
		cookie := cookieKit.NewCookie(name, strconv.Itoa(0), "", "", 3600, false, false, http.SameSiteDefaultMode)
		cookieKit.SetCookie(ctx.Writer, cookie)

		ctx.String(http.StatusOK, "set")
	})

	engine.Any("/add", func(ctx *gin.Context) {
		cookie, err := cookieKit.GetCookie(ctx.Request, name)
		if err != nil {
			ctx.String(http.StatusOK, err.Error())
			return
		}

		value, err := urlKit.DecodeURIComponent(cookie.Value)
		if err != nil {
			ctx.String(http.StatusOK, err.Error())
			return
		}
		i, err := strconv.Atoi(value)
		if err != nil {
			ctx.String(http.StatusOK, err.Error())
			return
		}
		i++
		cookie.Value = urlKit.EncodeURIComponent(strconv.Itoa(i))

		// 后端修改了cookie，如果不set回去，前端的cookie将不会同步修改
		cookieKit.SetCookie(ctx.Writer, cookie)

		ctx.String(http.StatusOK, "get "+cookie.Value)
	})

	engine.Any("/del", func(ctx *gin.Context) {
		cookieKit.DeleteCookieByName(ctx.Request, ctx.Writer, name)

		ctx.String(http.StatusOK, "del")
	})

	if err := engine.Run(":80"); err != nil {
		panic(err)
	}
}
