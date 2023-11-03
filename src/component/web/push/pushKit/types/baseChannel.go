package types

import (
	"github.com/richelieu-yang/chimera/v2/src/mutexKit"
)

// BaseChannel
/*
!!!:
(1) 此类型实现了部分 Channel 接口，
(2) 此类型的子类应当实现 Channel 接口（主要是: Push()、Close()、bind、unbind），不能覆盖父类的方法.
*/
type BaseChannel struct {
	RWMutex mutexKit.RWMutex

	Id    string
	Bsid  string
	User  string
	Group string
	Data  interface{}

	Closed    bool
	Listeners Listeners
}

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

func (channel *BaseChannel) GetId() string {
	return channel.Id
}

func (channel *BaseChannel) GetBsid() string {
	return channel.Bsid
}

func (channel *BaseChannel) ClearBsid() {
	channel.Bsid = ""
}

func (channel *BaseChannel) GetUser() string {
	return channel.User
}

func (channel *BaseChannel) ClearUser() {
	channel.User = ""
}

func (channel *BaseChannel) GetGroup() string {
	return channel.Group
}

func (channel *BaseChannel) ClearGroup() {
	channel.Group = ""
}

func (channel *BaseChannel) GetData() interface{} {
	return channel.Data
}

func (channel *BaseChannel) ClearData() {
	channel.Data = nil
}
