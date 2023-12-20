package freecacheKit

import (
	"github.com/coocood/freecache"
)

var NewCache func(size int) (cache *freecache.Cache) = freecache.NewCache
