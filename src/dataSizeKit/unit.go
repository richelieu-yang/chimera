package dataSizeKit

type Unit struct {
	value uint64
	str   string
}

func (su *Unit) GetValue() uint64 {
	if su == nil {
		return B.value
	}
	return su.value
}
func (su *Unit) String() string {
	if su == nil {
		return B.str
	}
	return su.str
}

var (
	B = &Unit{
		value: 1,
		str:   "B",
	}
	KB = &Unit{
		value: 1024,
		str:   "KB",
	}
	MB = &Unit{
		value: 1048576,
		str:   "MB",
	}
	GB = &Unit{
		value: 1073741824,
		str:   "GB",
	}
	TB = &Unit{
		value: 1099511627776,
		str:   "TB",
	}
)
