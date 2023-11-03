package types

import (
	"github.com/richelieu-yang/chimera/v2/src/mutexKit"
)

// BaseChannel
/*
!!!:
(1) 此类型实现了部分 Channel 接口，
(2) 此类型的子类应当实现 Channel 接口（主要是: Push()、Close()、bind、unbind），不能覆盖父类的方法.
	(由于unbind可能涉及Close()，因此只能在子类中实现)
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
	/* 读锁 */
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

	/* 写锁 */
	channel.RWMutex.LockFunc(func() {
		if channel.Closed {
			return
		}
		channel.Closed = true
		flag = true
	})
	return
}

func (channel *BaseChannel) GetId() (id string) {
	// 无需加解锁，因为: id是不变的
	return channel.Id
}

func (channel *BaseChannel) GetBsid() (bsid string) {
	/* 读锁 */
	channel.RWMutex.RLockFunc(func() {
		bsid = channel.Bsid
	})
	return
}

func (channel *BaseChannel) ClearBsid() {
	if channel.Bsid == "" {
		return
	}

	/* 写锁 */
	channel.RWMutex.LockFunc(func() {
		channel.Bsid = ""
	})
}

func (channel *BaseChannel) GetUser() (user string) {
	/* 读锁 */
	channel.RWMutex.RLockFunc(func() {
		user = channel.User
	})
	return
}

func (channel *BaseChannel) ClearUser() {
	if channel.User == "" {
		return
	}

	/* 写锁 */
	channel.RWMutex.LockFunc(func() {
		channel.User = ""
	})
}

func (channel *BaseChannel) GetGroup() (group string) {
	/* 读锁 */
	channel.RWMutex.RLockFunc(func() {
		group = channel.Group
	})
	return
}

func (channel *BaseChannel) ClearGroup() {
	if channel.Group == "" {
		return
	}

	/* 写锁 */
	channel.RWMutex.LockFunc(func() {
		channel.Group = ""
	})
}

func (channel *BaseChannel) GetData() (data interface{}) {
	/* 读锁 */
	channel.RWMutex.RLockFunc(func() {
		data = channel.Data
	})
	return
}

func (channel *BaseChannel) ClearData() {
	if channel.Data == nil {
		return
	}

	/* 写锁 */
	channel.RWMutex.LockFunc(func() {
		channel.Data = nil
	})
}
