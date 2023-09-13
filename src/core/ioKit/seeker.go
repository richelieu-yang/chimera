package ioKit

import "io"

// SeekToStart 回到最前面.
func SeekToStart(seeker io.Seeker) (int64, error) {
	return seeker.Seek(0, io.SeekStart)
}
