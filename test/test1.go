package main

import (
	"github.com/gogf/gf/v2/container/gmap"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gutil"
)

func main() {
	array := g.Slice{2, 3, 1, 5, 4, 6, 8, 7, 9}
	hashMap := gmap.New(true)
	listMap := gmap.NewListMap(true)
	treeMap := gmap.NewTreeMap(gutil.ComparatorInt, true)

}
