package wsKit

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/richelieu-yang/chimera/v2/src/component/web/push/pushKit/types"
	"github.com/richelieu-yang/chimera/v2/src/core/errorKit"
	"github.com/richelieu-yang/chimera/v2/src/core/strKit"
	"github.com/richelieu-yang/chimera/v2/src/mutexKit"
	"net/http"
	"time"
)

type WsProcessor struct {
	types.Processor

	// upgrader 是并发安全的
	upgrader *websocket.Upgrader

	idGenerator func() (string, error)
	listeners   types.Listeners
	msgType     messageType
}

func (p *WsProcessor) ProcessWithGin(ctx *gin.Context) {
	p.Process(ctx.Writer, ctx.Request)
}

func (p *WsProcessor) Process(w http.ResponseWriter, r *http.Request) {
	PolyfillWebSocketRequest(r)

	// 先判断是不是websocket请求
	if !IsWebSocketUpgrade(r) {
		failureInfo := "Not a websocket upgrade request"
		p.listeners.OnFailure(w, r, failureInfo)
		return
	}

	// Upgrade（升级为WebSocket协议）
	conn, err := p.upgrader.Upgrade(w, r, w.Header())
	if err != nil {
		failureInfo := fmt.Sprintf("Fail to upgrade because of error(%s)", err.Error())
		p.listeners.OnFailure(w, r, failureInfo)
		return
	}
	defer conn.Close()

	channel, err := p.newChannel(conn)
	if err != nil {
		err = errorKit.Wrap(err, "Fail to new channel")
		p.listeners.OnFailure(w, r, err.Error())
		return
	}

	p.listeners.OnHandshake(w, r, channel)

	conn.SetCloseHandler(func(code int, text string) error {
		if channel.SetClosed() {
			info := fmt.Sprintf("code: %d, text: %s", code, text)
			p.listeners.OnClose(channel, info)
		}

		// 默认的close handler
		message := websocket.FormatCloseMessage(code, text)
		_ = conn.WriteControl(websocket.CloseMessage, message, time.Now().Add(time.Second))
		return nil
	})

	/* 接收WebSocket客户端发来的消息 */
	for {
		messageType, data, err := conn.ReadMessage()
		if err != nil {
			if channel.SetClosed() {
				var closeErr *websocket.CloseError
				if errors.As(err, &closeErr) {
					info := fmt.Sprintf("code: %d, text: %s", closeErr.Code, closeErr.Text)
					p.listeners.OnClose(channel, info)
				} else {
					info := fmt.Sprintf("Fail to read message because of error(%s)", err.Error())
					p.listeners.OnClose(channel, info)
				}
			}
			break
		}
		p.listeners.OnMessage(channel, messageType, data)
	}
}

func (p *WsProcessor) newChannel(conn *websocket.Conn) (*WsChannel, error) {
	id, err := p.idGenerator()
	if err != nil {
		return nil, errorKit.Wrap(err, "Fail to generate id")
	}
	if err := strKit.AssertNotBlank(id, "id"); err != nil {
		return nil, err
	}

	channel := &WsChannel{
		BaseChannel: &types.BaseChannel{
			Id:        id,
			Bsid:      "",
			User:      "",
			Group:     "",
			RWMutex:   mutexKit.RWMutex{},
			Data:      nil,
			Closed:    false,
			Listeners: p.listeners,
		},
		conn:    conn,
		msgType: p.msgType,
	}
	return channel, nil
}
