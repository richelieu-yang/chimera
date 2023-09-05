package cmpKit

import "github.com/google/go-cmp/cmp"

var Equal func(x, y interface{}, opts ...cmp.Option) bool = cmp.Equal

var Diff func(x, y interface{}, opts ...cmp.Option) string = cmp.Diff
