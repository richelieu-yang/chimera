package otelKit

import (
	"go.opentelemetry.io/otel"
	"net/http"
	"testing"
)

func TestNewJaegerTracerProvider(t *testing.T) {
	tp, err := NewJaegerTracerProvider("http://localhost:14268/api/traces", "service", "environment", 1)
	if err != nil {
		panic(err)
	}
	otel.SetTracerProvider(tp)

	http.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		tracer := otel.Tracer("my-tracer")
		_, span := tracer.Start(ctx, "spanName")
		defer span.End()

		_, _ = w.Write([]byte("Hello, World!"))
	})
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
