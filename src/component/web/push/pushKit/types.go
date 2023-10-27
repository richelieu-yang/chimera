package pushKit

import (
	"github.com/richelieu-yang/chimera/v2/src/mutexKit"
)

type (
	Channel interface {
		// Push 推送（二进制）消息给客户端.
		Push(messageType MessageType, data []byte) error

		// Close 后端主动关闭通道.
		Close() error
	}

	BaseChannel struct {
		Id    string
		Bsid  string
		User  string
		Group string

		RWMutex mutexKit.RWMutex

		Data   interface{}
		Closed bool
	}
)

func (channel *BaseChannel) IsClosed() (rst bool) {
	// 读锁
	channel.RWMutex.RLockFunc(func() {
		rst = channel.Closed
	})
	return
}

// SetClosed
/*
@return true: 	设置成功
		false:	设置失败（因为已经被设置关闭）
*/
func (channel *BaseChannel) SetClosed() (flag bool) {
	if channel.Closed {
		return
	}

	// 写锁
	channel.RWMutex.LockFunc(func() {
		if channel.Closed {
			return
		}
		channel.Closed = true
		flag = true
	})
	return
}
