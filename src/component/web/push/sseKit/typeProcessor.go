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

func (p *SseProcessor) ProcessWithGin(ctx *gin.Context) {
	p.Process(ctx.Writer, ctx.Request)
}

func (p *SseProcessor) Process(w http.ResponseWriter, r *http.Request) {
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

	/*
		!!!: gin.Context.Done() 和 r.Context().Done() 不同
		(1) case为 w.(http.CloseNotifier).CloseNotify() 和 gin.Context.Done()，前端断开会走到 w.(http.CloseNotifier).CloseNotify()
		(2) case为 w.(http.CloseNotifier).CloseNotify() 和 r.Context().Done()，前端断开会走到 r.Context().Done()
	*/
	select {
	case <-r.Context().Done():
		p.listeners.OnClose(channel, "Context done")
	//case <-w.(http.CloseNotifier).CloseNotify():
	//	p.listeners.OnClose(channel, "Connection closed")
	case <-channel.closeCh:
		// 后端主动断开连接
		p.listeners.OnClose(channel, "Connection closed by backend")
	}
}

func (p *SseProcessor) newChannel(w http.ResponseWriter, r *http.Request) (*SseChannel, error) {
	id, err := p.idGenerator()
	if err != nil {
		return nil, errorKit.Wrap(err, "Fail to generate id")
	}
	if err := strKit.AssertNotBlank(id, "id"); err != nil {
		return nil, err
	}

	channel := &SseChannel{
		BaseChannel: &pushKit.BaseChannel{
			RWMutex:   mutexKit.RWMutex{},
			Id:        id,
			Bsid:      "",
			User:      "",
			Group:     "",
			Data:      nil,
			Closed:    false,
			Listeners: p.listeners,
		},
		w:       w,
		r:       r,
		msgType: p.msgType,
		closeCh: make(chan struct{}, 1),
	}
	return channel, nil
}
