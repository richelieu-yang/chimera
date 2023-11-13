package sseKit

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/richelieu-yang/chimera/v2/src/component/web/httpKit"
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

	closeCh := make(chan string, 1)
	channel, err := p.newChannel(w, r, closeCh)
	if err != nil {
		err = errorKit.Wrap(err, "Fail to new channel")
		p.listeners.OnFailure(w, r, err.Error())
		return
	}
	p.listeners.OnHandshake(w, r, channel)

	/*
		!!!: gin.Context.Done() 和 r.Context().Done() 不同，因为 gin.Context.Done() 返回nil（普通Gin Server情况下）.
		(1) case为 w.(http.CloseNotifier).CloseNotify() 和 gin.Context.Done()，前端断开会走到 w.(http.CloseNotifier).CloseNotify()
		(2) case为 w.(http.CloseNotifier).CloseNotify() 和 r.Context().Done()，前端断开会走到 r.Context().Done()
	*/
	select {
	case <-r.Context().Done():
		// 前端主动断开连接
		p.listeners.OnClose(channel, "Context done")
	//case <-w.(http.CloseNotifier).CloseNotify():
	//	p.listeners.OnClose(channel, "Connection closed")
	case reason := <-closeCh:
		// 后端主动断开连接
		closeInfo := fmt.Sprintf("Connection is closed by backend with reason(%s)", reason)
		p.listeners.OnClose(channel, closeInfo)
	}
}

func (p *SseProcessor) newChannel(w http.ResponseWriter, r *http.Request, closeCh chan string) (pushKit.Channel, error) {
	id, err := p.idGenerator()
	if err != nil {
		return nil, errorKit.Wrap(err, "Fail to generate id")
	}
	if err := strKit.AssertNotBlank(id, "id"); err != nil {
		return nil, err
	}

	ip, err := httpKit.GetClientIP(r)
	if err != nil {
		ip = err.Error()
	}

	channel := &SseChannel{
		BaseChannel: pushKit.BaseChannel{
			RWMutex:   mutexKit.RWMutex{},
			CloseCh:   closeCh,
			ClientIP:  ip,
			Type:      "SSE",
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
	}
	return channel, nil
}
