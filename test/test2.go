package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
)

func main() {
	router := gin.New()
	router.GET("/test1", func(c *gin.Context) {
		c.String(http.StatusOK, "this is 8889")
	})

	if err := router.Run(":8889"); err != nil {
		logrus.Fatal(err)
	}
}
