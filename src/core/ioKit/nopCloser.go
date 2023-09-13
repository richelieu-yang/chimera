package ioKit

import (
	"io"
)

// NopCloser
/*
PS: 返回值调用Close()将什么都不会做，直接返回nil.
*/
var NopCloser func(reader io.Reader) io.ReadCloser = io.NopCloser

// NopCloserToReadSeeker 给 io.ReadSeeker 实例加上Close()（虽然这个Close()什么都不做）.
func NopCloserToReadSeeker(readSeeker io.ReadSeeker) io.ReadSeekCloser {
	return &readSeekCloser{
		readSeeker: readSeeker,
	}
}

//type nopCloserToWriter struct {
//	io.Writer
//}
//
//func (nopCloserToWriter) Close() error {
//	return nil
//}
//
//type nopCloserToReadSeeker struct {
//	io.ReadSeeker
//}
//
//func (nopCloserToReadSeeker) Close() error {
//	return nil
//}
//
//// NopCloserToWriter
///*
//PS: 返回值调用Close()将什么都不会做，直接返回nil.
//*/
//func NopCloserToWriter(writer io.Writer) io.WriteCloser {
//	return &nopCloserToWriter{
//		writer,
//	}
//}
//
//func NopCloserToReadSeeker(readSeeker io.ReadSeeker) io.ReadSeekCloser {
//	return &nopCloserToReadSeeker{
//		ReadSeeker: readSeeker,
//	}
//}
