package types

import (
	"github.com/richelieu-yang/chimera/v2/src/component/web/push/pushKit"
	"github.com/richelieu-yang/chimera/v2/src/mutexKit"
)

// BaseChannel
/*
!!!: 此类型的子类应当:
	(1) 实现 Channel 接口（主要是bind、unbind方法），
	(2) 覆盖 Push、Close() 方法.
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

func (channel *BaseChannel) Push(data []byte) error {
	panic("implement me")
}

func (channel *BaseChannel) Close(reason string) error {
	panic("implement me")
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

func (channel *BaseChannel) BindGroup(group string) {
	pushKit.BindGroup(channel, group)
}

func (channel *BaseChannel) BindUser(user string) {
	pushKit.BindUser(channel, user)
}

func (channel *BaseChannel) BindBsid(bsid string) {
	pushKit.BindBsid(channel, bsid)
}

func (channel *BaseChannel) Unbind() {
	pushKit.UnBindId(channel, channel.Id)
	pushKit.UnbindBsid(channel, channel.Bsid)
	pushKit.UnbindUser(channel, channel.User)
	pushKit.UnbindGroup(channel, channel.Group)
}

func (channel *BaseChannel) GetId() string {
	return channel.Id
}

func (channel *BaseChannel) GetBsid() string {
	return channel.Bsid
}

func (channel *BaseChannel) GetUser() string {
	return channel.User
}

func (channel *BaseChannel) GetGroup() string {
	return channel.Group
}

func (channel *BaseChannel) GetData() interface{} {
	return channel.Data
}
