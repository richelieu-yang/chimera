package sseKit

import (
	"github.com/gin-gonic/gin"
	"github.com/richelieu-yang/chimera/v2/src/component/web/push/pushKit"
	"github.com/richelieu-yang/chimera/v2/src/core/errorKit"
	"github.com/richelieu-yang/chimera/v2/src/core/strKit"
	"github.com/richelieu-yang/chimera/v2/src/mutexKit"
	"net/http"
)

type SseProcessor struct {
	pushKit.Processor

	idGenerator func() (string, error)
	listeners   pushKit.Listeners
	msgType     messageType
}

func (p *SseProcessor) HandleWithGin(ctx *gin.Context) {
	p.Handle(ctx.Writer, ctx.Request)
}

func (p *SseProcessor) Handle(w http.ResponseWriter, r *http.Request) {
	if err := IsSseSupported(w, r); err != nil {
		p.listeners.OnFailure(w, r, err.Error())
		return
	}

	// 设置 response header
	SetHeaders(w)

	channel, err := p.newChannel(w, r)
	if err != nil {
		err = errorKit.Wrap(err, "Fail to new channel")
		p.listeners.OnFailure(w, r, err.Error())
		return
	}
	p.listeners.OnHandshake(w, r, channel)

	select {
	case <-r.Context().Done():
		p.listeners.OnClose(channel, "Context done")
	case <-w.(http.CloseNotifier).CloseNotify():
		// SSE客户端关闭后，会走此处
		p.listeners.OnClose(channel, "Connection closed")
	}
}

func (p *SseProcessor) newChannel(w http.ResponseWriter, r *http.Request) (pushKit.Channel, error) {
	id, err := p.idGenerator()
	if err != nil {
		return nil, errorKit.Wrap(err, "Fail to generate id")
	}
	if err := strKit.AssertNotBlank(id, "id"); err != nil {
		return nil, err
	}

	channel := &SseChannel{
		BaseChannel: &pushKit.BaseChannel{
			Id:        id,
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
	return channel, nil
}
