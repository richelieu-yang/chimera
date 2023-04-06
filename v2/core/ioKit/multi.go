package ioKit

import (
	"io"
)

func MultiReader(readers ...io.Reader) io.Reader {
	return io.MultiReader(readers...)
}

func MultiWriter(writers ...io.Writer) io.Writer {
	return io.MultiWriter(writers...)
}
