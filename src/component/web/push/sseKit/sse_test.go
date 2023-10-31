package sseKit

import (
	"github.com/gin-gonic/gin"
	"github.com/richelieu-yang/chimera/v2/src/log/logrusKit"
	"github.com/sirupsen/logrus"
	"testing"
)

func TestNewProcessor(t *testing.T) {
	logrusKit.MustSetUp(nil)
	processor, err := NewProcessor()

	engine := gin.Default()

	if err := engine.Run(":80"); err != nil {
		logrus.Fatal(err)
	}
}
