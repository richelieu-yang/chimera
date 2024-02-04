package otelKit

import (
	"context"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/baggage"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/trace"
	"net/http"
)

func GetRemoteSpanCtx(r *http.Request) (remoteSpanCtx context.Context, err error) {
	//var propagator = propagation.TextMapPropagator(propagation.Baggage{})
	//propagatorCtx := propagator.Extract(r.Context(), propagation.HeaderCarrier(r.Header))
	propagatorCtx := otel.GetTextMapPropagator().Extract(r.Context(), propagation.HeaderCarrier(r.Header))

	b := baggage.FromContext(propagatorCtx)
	traceId, err := trace.TraceIDFromHex(b.Member("trace-id").Value())
	if err != nil {
		return
	}
	spanId, err := trace.SpanIDFromHex(b.Member("span-id").Value())
	if err != nil {
		return
	}
	spanCtx := trace.NewSpanContext(trace.SpanContextConfig{
		TraceID:    traceId,
		SpanID:     spanId,
		TraceFlags: trace.FlagsSampled, //这个没写，是不会记录的
		TraceState: trace.TraceState{},
		Remote:     true,
	})
	remoteSpanCtx = trace.ContextWithRemoteSpanContext(propagatorCtx, spanCtx)
	return
}
