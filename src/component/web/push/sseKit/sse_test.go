package sseKit

import (
	"github.com/gin-gonic/gin"
	"github.com/richelieu-yang/chimera/v2/src/component/web/ginKit"
	"github.com/richelieu-yang/chimera/v2/src/component/web/push/pushKit"
	"github.com/richelieu-yang/chimera/v2/src/log/logrusKit"
	"github.com/sirupsen/logrus"
	"net/http"
	"testing"
)

type listener struct {
	pushKit.Listener
}

func (l *listener) OnFailure(w http.ResponseWriter, r *http.Request, failureInfo string) {
	logrus.WithField("failureInfo", failureInfo).Error("OnFailure")
}

func (l *listener) OnHandshake(w http.ResponseWriter, r *http.Request, channel pushKit.Channel) {
	logrus.Info("OnHandshake")

	if err := channel.Push([]byte("Hello!")); err != nil {
		logrus.Error(err)
	}
}

func (l *listener) OnMessage(channel pushKit.Channel, messageType int, data []byte) {
	logrus.Info("OnMessage")
}

func (l *listener) OnClose(channel pushKit.Channel, closeInfo string) {
	logrus.WithField("closeInfo", closeInfo).Info("OnClose")
}

func TestNewProcessor(t *testing.T) {
	logrusKit.MustSetUp(nil)
	processor, err := NewProcessor(nil, &listener{}, MessageTypeRaw)
	if err != nil {
		panic(err)
	}

	engine := gin.Default()
	engine.Use(ginKit.NewCorsMiddleware(nil))
	engine.GET("/sse", processor.ProcessWithGin)
	if err := engine.Run(":80"); err != nil {
		panic(err)
	}
}
