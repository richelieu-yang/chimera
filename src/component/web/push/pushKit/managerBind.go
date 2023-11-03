package pushKit

import (
	"fmt"
	"github.com/richelieu-yang/chimera/v2/src/component/web/push/pushKit/types"
	"github.com/richelieu-yang/chimera/v2/src/core/setKit"
	"github.com/richelieu-yang/chimera/v2/src/core/strKit"
)

func BindId(channel types.Channel, id string) {
	if strKit.IsEmpty(id) {
		return
	}

	// 写锁
	idMap.RWLock.LockFunc(func() {
		if old, ok := bsidMap.Map[id]; ok {
			_ = old.Close(fmt.Sprintf("id(%s) is replaced by new channel", id))
		}
		idMap.Map[id] = channel
	})
}

func BindBsid(channel types.Channel, bsid string) {
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

func BindUser(channel types.Channel, user string) {
	if strKit.IsEmpty(user) {
		return
	}

	var userSet *setKit.SetWithLock[types.Channel]
	// 写锁
	userMap.RWLock.LockFunc(func() {
		userSet = userMap.Map[user]
		if userSet == nil {
			userSet = setKit.NewSetWithLock[types.Channel]()
			userMap.Map[user] = userSet
		}
	})

	// 写锁
	userSet.RWLock.LockFunc(func() {
		_ = userSet.Set.Add(channel)
	})
}

func BindGroup(channel types.Channel, group string) {
	if strKit.IsEmpty(group) {
		return
	}

	// 写锁
	var groupSet *setKit.SetWithLock[types.Channel]
	groupMap.RWLock.LockFunc(func() {
		groupSet = userMap.Map[group]
		if groupSet == nil {
			groupSet = setKit.NewSetWithLock[types.Channel]()
			groupMap.Map[group] = groupSet
		}
	})

	// 写锁
	groupSet.RWLock.LockFunc(func() {
		_ = groupSet.Set.Add(channel)
	})
}
