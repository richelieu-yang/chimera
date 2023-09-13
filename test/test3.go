package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/events", func(w http.ResponseWriter, r *http.Request) {
		flusher, ok := w.(http.Flusher)
		if !ok {
			http.Error(w, "Streaming unsupported!", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "text/event-stream")
		w.Header().Set("Cache-Control", "no-cache")
		w.Header().Set("Connection", "keep-alive")

		for i := 1; i <= 10; i++ {
			fmt.Fprintf(w, "data: Message %d\n\n", i)
			flusher.Flush()
		}
	})

	http.ListenAndServe(":8080", nil)
}
