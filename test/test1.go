package main

import (
	"github.com/gin-gonic/gin"
	"github.com/richelieu42/chimera/src/core/strKit"
	"github.com/richelieu42/chimera/src/http/httpKit"
	"github.com/sirupsen/logrus"
	"net/http"
)

func main() {
	router := gin.New()
	router.GET("/test", func(c *gin.Context) {
		err := httpKit.Proxy(c.Writer, c.Request, nil, "http", "127.0.0.1:80", strKit.GetStringPtr("/ws/connect"), nil)
		if err != nil {
			c.String(http.StatusOK, err.Error())
		}
	})

	// 方法1
	if err := router.RunTLS("0.0.0.0:8888", "ssl.pem", "ssl.key"); err != nil {
		logrus.Panic(err)
	}
	// 方法2
	//server := &http.Server{
	//	Addr:    "0.0.0.0:10679",
	//	Handler: router,
	//}
	//_ = server.ListenAndServeTLS("./certs/server.cer", "./certs/server.key")
}
