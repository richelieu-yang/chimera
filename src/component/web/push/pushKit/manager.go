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

	// groupMap
	/*
		key: group（一对多）
	*/
	groupMap = mapKit.NewMapWithLock[string, setKit.SetWithLock[Channel]]()

	// userMap
	/*
		key: user（一对多）
	*/
	userMap = mapKit.NewMapWithLock[string, setKit.SetWithLock[Channel]]()

	// bsidMap
	/*
		key: bsid（一对一）
	*/
	bsidMap = mapKit.NewMapWithLock[string, Channel]()
)
