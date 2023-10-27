package types

import (
	"errors"
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/richelieu-yang/chimera/v2/src/component/web/push/pushKit/types"
	"github.com/richelieu-yang/chimera/v2/src/component/web/push/wsKit"
	"github.com/richelieu-yang/chimera/v2/src/core/errorKit"
	"github.com/richelieu-yang/chimera/v2/src/core/interfaceKit"
	"github.com/richelieu-yang/chimera/v2/src/core/strKit"
	"github.com/richelieu-yang/chimera/v2/src/core/timeKit"
	"github.com/richelieu-yang/chimera/v2/src/idKit"
	"github.com/richelieu-yang/chimera/v2/src/mutexKit"
	"net/http"
	"time"
)

type Processor struct {
	// upgrader 是并发安全的
	upgrader *websocket.Upgrader

	idGenerator func() (string, error)

	listener types.Listener
}

func (p *Processor) NewChannel(conn *websocket.Conn) (*WsChannel, error) {
	id, err := p.idGenerator()
	if err != nil {
		return nil, err
	}
	if err := strKit.AssertNotBlank(id, "id"); err != nil {
		return nil, err
	}

	return &WsChannel{
		BaseChannel: types.BaseChannel{
			Id:       id,
			Bsid:     "",
			User:     "",
			Group:    "",
			RWMutex:  mutexKit.RWMutex{},
			Data:     nil,
			Closed:   false,
			Listener: p.listener,
		},
		conn: conn,
	}, nil
}

func (p *Processor) Handle(w http.ResponseWriter, r *http.Request) {
	wsKit.PolyfillWebSocketRequest(r)

	// 先判断是不是websocket请求
	if !websocket.IsWebSocketUpgrade(r) {
		failureInfo := "Not a websocket upgrade request"
		p.listener.OnFailure(w, r, failureInfo)
		return
	}

	// Upgrade（升级为WebSocket协议）
	conn, err := p.upgrader.Upgrade(w, r, w.Header())
	if err != nil {
		failureInfo := fmt.Sprintf("Fail to upgrade because of error(%s)", err.Error())
		p.listener.OnFailure(w, r, failureInfo)
		return
	}
	defer func() {
		_ = conn.Close()
	}()

	channel, err := p.NewChannel(conn)
	if err != nil {
		err = errorKit.Wrap(err, "Fail to new channel")
		p.listener.OnFailure(w, r, err.Error())
		return
	}

	p.listener.OnHandshake(w, r, channel)

	conn.SetCloseHandler(func(code int, text string) error {
		if channel.SetClosed() {
			info := fmt.Sprintf("code: %d, text: %s", code, text)
			p.listener.OnClose(channel, info)
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
					p.listener.OnClose(channel, info)
				} else {
					info := fmt.Sprintf("Fail to read message because of error(%s)", err.Error())
					p.listener.OnClose(channel, info)
				}
			}
			break
		}
		p.listener.OnMessage(channel, messageType, data)
	}
}

// NewProcessor
/*
@param handshakeTimeout	默认3s
@param checkOrigin		可以为nil（允许所有）
@param idGenerator		可以为nil（使用xid）
@param listener			不能为nil
*/
func NewProcessor(handshakeTimeout time.Duration, checkOrigin func(r *http.Request) bool, idGenerator func() (string, error), listener types.Listener) (*Processor, error) {
	handshakeTimeout = timeKit.ToDefaultDurationIfInvalid(handshakeTimeout, time.Second*3)
	if checkOrigin == nil {
		checkOrigin = func(r *http.Request) bool {
			// 允许跨域
			return true
		}
	}
	upgrader := &websocket.Upgrader{
		HandshakeTimeout: handshakeTimeout,
		CheckOrigin: func(r *http.Request) bool {
			// 允许跨域
			return true
		},
	}

	if idGenerator == nil {
		idGenerator = func() (string, error) {
			return idKit.NewXid(), nil
		}
	}

	if err := interfaceKit.AssertNotNil(listener, "listener"); err != nil {
		return nil, err
	}

	return &Processor{
		upgrader:    upgrader,
		idGenerator: idGenerator,
		listener:    listener,
	}, nil
}
