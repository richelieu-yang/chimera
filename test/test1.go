package main

import (
	"github.com/richelieu42/go-scales/src/core/setKit"
)

type Bean struct {
	Id int
}

func main() {
	b0 := &Bean{Id: 0}
	b1 := &Bean{Id: 0}
	b2 := &Bean{Id: 0}

	set := setKit.NewSet(false, b0, b1)
	set.Add(b0)
	set.Add(b1)
	set.Add(b2)
	set.Remove(b1)
}
