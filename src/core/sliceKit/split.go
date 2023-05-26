package sliceKit

// Split 分割 []T
func Split[T any](buf []T, limit int) [][]T {
	var chunk []T
	chunks := make([][]T, 0, len(buf)/limit+1)

	for len(buf) >= limit {
		chunk, buf = buf[:limit], buf[limit:]
		chunks = append(chunks, chunk)
	}
	if len(buf) > 0 {
		chunks = append(chunks, buf[:])
	}
	return chunks
}
