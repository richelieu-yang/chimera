package dataSizeKit

type sizeUnit struct {
	value uint64
	str   string
}

func (su *sizeUnit) GetValue() uint64 {
	if su == nil {
		return B.value
	}
	return su.value
}
func (su *sizeUnit) String() string {
	if su == nil {
		return B.str
	}
	return su.str
}

var (
	B = &sizeUnit{
		value: 1,
		str:   "B",
	}
	KB = &sizeUnit{
		value: 1024,
		str:   "KB",
	}
	MB = &sizeUnit{
		value: 1048576,
		str:   "MB",
	}
	GB = &sizeUnit{
		value: 1073741824,
		str:   "GB",
	}
	TB = &sizeUnit{
		value: 1099511627776,
		str:   "TB",
	}
)
