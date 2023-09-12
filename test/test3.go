package main

import (
	"fmt"
	"github.com/richelieu-yang/chimera/v2/src/web/httpClientKit"
)

func main() {
	fmt.Println(httpClientKit.Get("http://127.0.0.1/c?a=%E6%B5%8B%E8%AF%95&b=b&c=c&d=%E5%A4%A7#%E4%B8%AD%E9%83%A8"))

	//http.HandleFunc("/events", func(w http.ResponseWriter, r *http.Request) {
	//	flusher, ok := w.(http.Flusher)
	//	if !ok {
	//		http.Error(w, "Streaming unsupported!", http.StatusInternalServerError)
	//		return
	//	}
	//
	//	w.Header().Set("Content-Type", "text/event-stream")
	//	w.Header().Set("Cache-Control", "no-cache")
	//	w.Header().Set("Connection", "keep-alive")
	//
	//	for i := 1; i <= 10; i++ {
	//		fmt.Fprintf(w, "data: Message %d\n\n", i)
	//		flusher.Flush()
	//	}
	//})
	//
	//http.ListenAndServe(":8080", nil)
}
