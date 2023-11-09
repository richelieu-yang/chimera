package pushKit

import (
	"github.com/richelieu-yang/chimera/v2/src/core/mapKit"
	"github.com/richelieu-yang/chimera/v2/src/core/setKit"
)

var (
	// idMap 即allMap
	/*
		key: id（一对一）
	*/
	idMap = mapKit.NewMapWithLock[string, Channel]()

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

// GetChannelByBsid （读锁）根据 bsid 获取Channel.
/*
@return 可能为nil
*/
func GetChannelByBsid(bsid string) (channel Channel) {
	/* 读锁 */
	bsidMap.RLockFunc(func() {
		channel = bsidMap.Map[bsid]
	})
	return
}

// GetUserSet （读锁）获取 userSet.
/*
@return 可能为nil
*/
func GetUserSet(user string) *setKit.SetWithLock[Channel] {
	var userSet *setKit.SetWithLock[Channel]

	/* 读锁 */
	userMap.RLockFunc(func() {
		userSet = userMap.Map[user]
	})
	return userSet
}

// GetGroupSet （读锁）获取 groupSet.
/*
@return 可能为nil
*/
func GetGroupSet(group string) *setKit.SetWithLock[Channel] {
	var groupSet *setKit.SetWithLock[Channel]

	/* 读锁 */
	groupMap.RLockFunc(func() {
		groupSet = userMap.Map[group]
	})
	return groupSet
}
