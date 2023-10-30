package sseKit

import (
	"net/http"
)

// IsSseSupported
/*
@return 为"": 支持SSE
*/
func IsSseSupported(w http.ResponseWriter, r *http.Request) (errText string) {
	if _, ok := w.(http.Flusher); !ok {
		// 不支持: 流信息（streaming）
		errText = "http.Flusher(Streaming) isn't supported"
		return
	}
	if _, ok := w.(http.CloseNotifier); !ok {
		// 不支持: 监听关闭
		errText = "http.CloseNotifier isn't supported!"
		return
	}
	return
}

// SetHeaders 设置response header.
func SetHeaders(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")
}
