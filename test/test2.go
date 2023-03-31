package main

import (
	"github.com/gin-gonic/gin"
	"github.com/unrolled/secure"
)

var (
	// certFile 证书文件
	certFile = "ssl.pem"
	// keyFile key文件
	keyFile = "ssl.key"
)

func main() {
	engine := gin.Default()
	engine.Use(TlsHandler())
	engine.GET("/test", func(ctx *gin.Context) {
		ctx.String(200, "hello [http 8000]")
	})

	engine1 := gin.Default()
	engine1.GET("/test", func(ctx *gin.Context) {
		ctx.String(200, "hello [https 8001]")
	})

	go func() {
		if err := engine.Run(":8000"); err != nil {
			panic(err)
		}
	}()
	go func() {
		if err := engine1.RunTLS(":8001", certFile, keyFile); err != nil {
			panic(err)
		}
	}()
	select {}
}

func TlsHandler() gin.HandlerFunc {
	secureMiddleware := secure.New(secure.Options{
		SSLRedirect: true,
		SSLHost:     "127.0.0.1:8001",
	})
	return func(c *gin.Context) {
		err := secureMiddleware.Process(c.Writer, c.Request)
		// If there was an error, do not continue.
		if err != nil {
			return
		}
		c.Next()
	}
}
