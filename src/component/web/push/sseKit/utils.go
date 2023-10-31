package sseKit

import (
	"github.com/richelieu-yang/chimera/v2/src/core/errorKit"
	"net/http"
)

// IsSseSupported
/*
@return 为"": 支持SSE
*/
func IsSseSupported(w http.ResponseWriter, r *http.Request) error {
	if _, ok := w.(http.Flusher); !ok {
		// 不支持: 流信息（streaming）
		return errorKit.New("http.Flusher(Streaming) isn't supported")
	}
	//if _, ok := w.(http.CloseNotifier); !ok {
	//	// 不支持: 监听关闭
	//	return errorKit.New("http.CloseNotifier isn't supported!")
	//}
	return nil
}

// SetHeaders 设置response header.
func SetHeaders(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")
}
