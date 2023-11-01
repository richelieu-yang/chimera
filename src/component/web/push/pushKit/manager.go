package pushKit

import (
	"github.com/richelieu-yang/chimera/v2/src/core/mapKit"
	"github.com/richelieu-yang/chimera/v2/src/core/setKit"
)

var (
	// allMap
	/*
		key: id（一对一）
	*/
	allMap = mapKit.NewMapWithLock[string, Channel]()

	// bsidMap
	/*
		key: bsid（一对一）
	*/
	bsidMap = mapKit.NewMapWithLock[string, Channel]()

	// userMap
	/*
		key: user（一对多）
	*/
	userMap = mapKit.NewMapWithLock[string, *setKit.SetWithLock[Channel]]()

	// groupMap
	/*
		key: group（一对多）
	*/
	groupMap = mapKit.NewMapWithLock[string, *setKit.SetWithLock[Channel]]()
)

func GetUserSet(user string) *setKit.SetWithLock[Channel] {
	var userSet *setKit.SetWithLock[Channel]
	userMap.RWLock.LockFunc(func() {
		userSet = userMap.Map[user]
	})
	return userSet
}

func GetGroupSet(group string) *setKit.SetWithLock[Channel] {
	var groupSet *setKit.SetWithLock[Channel]
	groupMap.RWLock.LockFunc(func() {
		groupSet = userMap.Map[group]
	})
	return groupSet
}
