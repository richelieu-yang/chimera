package intKit

import "github.com/spf13/cast"

var (
	ToUint func(i interface{}) uint = cast.ToUint

	ToUintE func(i interface{}) (uint, error) = cast.ToUintE

	ToUint8 func(i interface{}) uint8 = cast.ToUint8

	ToUint8E func(i interface{}) (uint8, error) = cast.ToUint8E

	ToUint16 func(i interface{}) uint16 = cast.ToUint16

	ToUint16E func(i interface{}) (uint16, error) = cast.ToUint16E

	ToUint32 func(i interface{}) uint32 = cast.ToUint32

	ToUint32E func(i interface{}) (uint32, error) = cast.ToUint32E

	ToUint64 func(i interface{}) uint64 = cast.ToUint64

	ToUint64E func(i interface{}) (uint64, error) = cast.ToUint64E
)
