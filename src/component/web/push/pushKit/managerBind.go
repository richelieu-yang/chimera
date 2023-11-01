package pushKit

import "github.com/richelieu-yang/chimera/v2/src/core/strKit"

func Bind(channel Channel, group, user, bsid string) {
	if strKit.IsNotEmpty(group) {

	}
	if strKit.IsNotEmpty(user) {

	}
	if strKit.IsNotEmpty(bsid) {

	}

}

func BindGroup(channel Channel, group string) {
	if strKit.IsEmpty(group) {
		return
	}
}

func BindUser(channel Channel, user string) {
	if strKit.IsEmpty(user) {
		return
	}

}

func BindBsid(channel Channel, bsid string) {
	if strKit.IsEmpty(bsid) {
		return
	}

	// 写锁
	bsidMap.RWLock.LockFunc(func() {
		if old, ok := bsidMap.Map[bsid]; ok {
			_ = old.Close()
		}

		bsidMap.Map[bsid] = channel
	})
}
