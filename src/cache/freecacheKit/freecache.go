package freecacheKit

import (
	"github.com/coocood/freecache"
)

// NewCache
/*
@param size 单位: byte
*/
var NewCache func(size int) (cache *freecache.Cache) = freecache.NewCache
