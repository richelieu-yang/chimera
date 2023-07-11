package jsonKit

import (
	"github.com/bytedance/sonic"
	"io"
)

// NewEncoder 编码器（to json）
func NewEncoder(api sonic.API, writer io.Writer) sonic.Encoder {
	return api.NewEncoder(writer)
}
