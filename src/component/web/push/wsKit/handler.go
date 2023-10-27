package wsKit

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/richelieu-yang/chimera/v2/src/component/web/push/pushKit/types"
	"github.com/richelieu-yang/chimera/v2/src/core/errorKit"
	"github.com/richelieu-yang/chimera/v2/src/core/interfaceKit"
	"net/http"
	"time"
)

// upgrader 并发安全的
var upgrader = websocket.Upgrader{
	HandshakeTimeout: time.Second * 3,
	CheckOrigin: func(r *http.Request) bool {
		// 允许跨域
		return true
	},
}

// NewGinHandler
/*
@param listener 不能为nil
*/
func NewGinHandler(listener types.Listener) (gin.HandlerFunc, error) {
	httpHandler, err := NewHttpHandler(listener)
	if err != nil {
		return nil, err
	}

	return func(ctx *gin.Context) {
		httpHandler(ctx.Writer, ctx.Request)
	}, nil
}

// NewHttpHandler
/*
@param listener 不能为nil
*/
func NewHttpHandler(listener types.Listener) (func(w http.ResponseWriter, r *http.Request), error) {
	if err := interfaceKit.AssertNotNil(listener, "Listener"); err != nil {
		return nil, err
	}

	return func(w http.ResponseWriter, r *http.Request) {
		PolyfillWebSocketRequest(r)

		// 先判断是不是websocket请求
		if !websocket.IsWebSocketUpgrade(r) {
			listener.OnFailure(w, r, "not WebSocket request")
			return
		}

		// Upgrade（升级为WebSocket协议）
		conn, err := upgrader.Upgrade(w, r, w.Header())
		if err != nil {
			err = errorKit.Wrap(err, "Fail to upgrade")
			listener.OnFailure(w, r, err.Error())
			return
		}
		defer conn.Close()

		channel := NewChannel(conn, listener)
		/* 监听: WebSocket客户端主动关闭连接 */
		conn.SetCloseHandler(func(code int, text string) error {
			channel.SetClosed()

			if RemoveChannel(channel) {
				channel.GetListener().OnCloseByFrontend(channel, code, text)
			}

			// 默认的close handler
			message := websocket.FormatCloseMessage(code, text)
			_ = conn.WriteControl(websocket.CloseMessage, message, time.Now().Add(time.Second))
			return nil
		})
		AddChannel(channel)
		listener.OnHandshake(channel, r)

		///* 绑定数据（通过url参数，有的话） */
		//bsid := httpKit.GetUrlParam(r, KeyBsid)
		//user := httpKit.GetUrlParam(r, KeyUser)
		//group := httpKit.GetUrlParam(r, KeyGroup)
		//if !strKit.IsAllEmpty(bsid, user, group) {
		//	channel.BindData(bsid, user, group)
		//}

		/* 接收WebSocket客户端发来的消息 */
		for {
			messageType, data, err := conn.ReadMessage()
			if err != nil {
				channel.SetClosed()

				if RemoveChannel(channel) {
					var closeErr *websocket.CloseError
					if errors.As(err, &closeErr) {
						listener.OnCloseByFrontend(channel, closeErr.Code, closeErr.Text)
					} else {
						listener.OnCloseByBackend(channel)
					}
				}
				break
			}
			listener.OnMessage(channel, messageType, data)
		}
	}, nil
}
