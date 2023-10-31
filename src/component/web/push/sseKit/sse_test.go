package sseKit

import (
	"github.com/gin-gonic/gin"
	"github.com/richelieu-yang/chimera/v2/src/component/web/push/pushKit"
	"github.com/richelieu-yang/chimera/v2/src/log/logrusKit"
	"net/http"
	"testing"
)

type listener struct {
	pushKit.Listener
}

func (l *listener) OnFailure(w http.ResponseWriter, r *http.Request, failureInfo string) {

}

func (l *listener) OnHandshake(w http.ResponseWriter, r *http.Request, channel pushKit.Channel) {

}

func (l *listener) OnMessage(channel pushKit.Channel, messageType int, data []byte) {

}

func (l *listener) OnClose(channel pushKit.Channel, closeInfo string) {

}

func TestNewProcessor(t *testing.T) {
	logrusKit.MustSetUp(nil)
	processor, err := NewProcessor(nil, &listener{}, MessageTypeRaw)
	if err != nil {
		panic(err)
	}

	engine := gin.Default()
	engine.GET("/sse", processor.ProcessWithGin)
	if err := engine.Run(":80"); err != nil {
		panic(err)
	}
}
