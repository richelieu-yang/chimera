package main

import (
	"github.com/gin-gonic/gin"
	"github.com/richelieu-yang/chimera/v2/src/json/jsonKit"
	"github.com/richelieu-yang/chimera/v2/src/log/logrusKit"
	"github.com/sirupsen/logrus"
	"net/http"
)

func init() {
	logrusKit.MustSetUp(nil)
}

// Param 定义待绑定的JSON结构体
type Param struct {
	Name  string   `json:"name"`
	Age   int      `json:"age"`
	Likes []string `json:"likes"`
}

func main() {
	engine := gin.Default()

	engine.POST("/test", func(ctx *gin.Context) {
		param := &Param{}

		if err := ctx.BindJSON(&param); err != nil {
			logrus.WithError(err).Error("fail to bind")
			return
		}

		logrusKit.DisableQuoteTemporarily(nil, func(logger *logrus.Logger) {
			json, _ := jsonKit.MarshalIndentToString(param, "", "    ")
			logger.Info(json)
		})
		ctx.String(http.StatusOK, "ok")
	})

	if err := engine.Run(":8888"); err != nil {
		panic(err)
	}
}
