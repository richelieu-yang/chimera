package sseKit

import (
	"github.com/gin-gonic/gin"
	"github.com/richelieu-yang/chimera/v2/src/component/web/push/pushKit"
	"github.com/richelieu-yang/chimera/v2/src/mutexKit"
	"net/http"
)

type SseProcessor struct {
	msgType   messageType
	listeners pushKit.Listeners
}

func (p *SseProcessor) HandleWithGin(ctx *gin.Context) {
	p.Handle(ctx.Writer, ctx.Request)
}

func (p *SseProcessor) Handle(w http.ResponseWriter, r *http.Request) {
	if errText := IsSseSupported(w, r); errText != "" {
		p.listeners.OnFailure(w, r, errText)
		return
	}

	// 设置 response header
	SetHeaders(w)

	channel := p.newChannel(w, r)
	p.listeners.OnHandshake(w, r, channel)

	select {
	case <-r.Context().Done():
		p.listeners.OnClose(channel, "Context done")
	case <-w.(http.CloseNotifier).CloseNotify():
		// SSE客户端关闭后，会走此处
		p.listeners.OnClose(channel, "Connection closed")
	}
}

func (p *SseProcessor) newChannel(w http.ResponseWriter, r *http.Request) pushKit.Channel {
	return &SseChannel{
		BaseChannel: &pushKit.BaseChannel{
			Id:        "",
			Bsid:      "",
			User:      "",
			Group:     "",
			RWMutex:   mutexKit.RWMutex{},
			Data:      nil,
			Closed:    false,
			Listeners: p.listeners,
		},
		w: w,
		r: r,
	}
}
