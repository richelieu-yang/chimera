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
	userMap = mapKit.NewMapWithLock[string, setKit.SetWithLock[Channel]]()

	// groupMap
	/*
		key: group（一对多）
	*/
	groupMap = mapKit.NewMapWithLock[string, setKit.SetWithLock[Channel]]()
)
