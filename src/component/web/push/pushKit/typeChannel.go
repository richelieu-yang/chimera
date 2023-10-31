package pushKit

import (
	"github.com/richelieu-yang/chimera/v2/src/mutexKit"
)

type (
	Channel interface {
		// Close 后端主动关闭通道.
		Close() error

		Push(data []byte) error

		IsClosed() (rst bool)

		SetClosed() (flag bool)
	}

	BaseChannel struct {
		RWMutex mutexKit.RWMutex

		Id    string
		Bsid  string
		User  string
		Group string
		Data  interface{}

		Closed    bool
		Listeners Listeners
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
PS: 如果返回值为true，应当触发 listener.onClose().

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
