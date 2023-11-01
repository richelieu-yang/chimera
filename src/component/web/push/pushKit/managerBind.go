package pushKit

import (
	"fmt"
	"github.com/richelieu-yang/chimera/v2/src/core/setKit"
	"github.com/richelieu-yang/chimera/v2/src/core/strKit"
)

func Bind(channel Channel, group, user, bsid string) {
	if strKit.IsNotEmpty(group) {

	}
	if strKit.IsNotEmpty(user) {

	}
	if strKit.IsNotEmpty(bsid) {

	}

}

func BindId(channel Channel, id string) {
	if strKit.IsEmpty(id) {
		return
	}

	// 写锁
	allMap.RWLock.LockFunc(func() {
		if old, ok := bsidMap.Map[id]; ok {
			_ = old.Close(fmt.Sprintf("id(%s) is replaced by new channel", id))
		}
		allMap.Map[id] = channel
	})
}

func BindBsid(channel Channel, bsid string) {
	if strKit.IsEmpty(bsid) {
		return
	}

	// 写锁
	bsidMap.RWLock.LockFunc(func() {
		if old, ok := bsidMap.Map[bsid]; ok {
			_ = old.Close(fmt.Sprintf("bsid(%s) is replaced by new channel", bsid))
		}
		bsidMap.Map[bsid] = channel
	})
}

func BindUser(channel Channel, user string) {
	if strKit.IsEmpty(user) {
		return
	}

	var userSet *setKit.SetWithLock[Channel]
	userMap.RWLock.LockFunc(func() {
		userSet = userMap.Map[user]
		if userSet == nil {
			userSet = setKit.NewSetWithLock[Channel]()
			userMap.Map[user] = userSet
		}
	})

	userSet.RWLock.LockFunc(func() {
		userSet.Set.Add(channel)
	})
}

func BindGroup(channel Channel, group string) {
	if strKit.IsEmpty(group) {
		return
	}
}
