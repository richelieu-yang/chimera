package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func main() {
	router := gin.New()
	router.GET("/ws/connect", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "success",
		})
	})
	// 可以直接用
	if err := router.RunTLS("0.0.0.0:8888", "ssl.pem", "ssl.key"); err != nil {
		logrus.Panic(err)
	}
	//server := &http.Server{Addr: "0.0.0.0:10679", Handler: router}
	//_ = server.ListenAndServeTLS("./certs/server.cer", "./certs/server.key")
}
