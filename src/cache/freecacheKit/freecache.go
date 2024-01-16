package freecacheKit

import (
	"github.com/coocood/freecache"
)

// NewCache
/*
@param size (1) 单位: byte
			(2) 最小512KB(freecache.minBufSize）
*/
var NewCache func(size int) (cache *freecache.Cache) = freecache.NewCache
