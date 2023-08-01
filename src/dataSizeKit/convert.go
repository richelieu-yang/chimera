package dataSizeKit

// ByteToMiB
/*
PS: IEC标准.
*/
func ByteToMiB(bytes uint64) float64 {
	return float64(bytes) / float64(MiB)
}

// ByteToGiB
/*
PS: IEC标准.
*/
func ByteToGiB(bytes uint64) float64 {
	return float64(bytes) / float64(GiB)
}
