package wsKit

type (
	Listener interface {
		/*
			连接握手成功后，会调用此方法.
		*/
		OnAfterHandshake(c *Channel)

		/*
			接收到前端发来的消息后，会调用此方法.
		*/
		OnMessage(c *Channel, msgType int, msgData []byte)

		/*
			连接断开后，会调用此方法.

			PS: 后端调用 Conn.Close() 并不会触发.
		*/
		OnAfterClose(c *Channel, code int, text string)
	}
)

var listener Listener = nil

func setListener(listener1 Listener) {
	listener = listener1
}
