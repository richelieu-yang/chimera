package otelKit

import (
	"context"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/trace"
)

// NewRemoteContext TODO:
/*
@param traceIdStr	请求头中，"trace-id"对应的值
@param spanIdStr	请求头中，"span-id"对应的值
*/
func NewRemoteContext(traceIdStr, spanIdStr string) (context.Context, error) {
	carrier := propagation.HeaderCarrier{}
	carrier.Set("trace-id", traceIdStr)

	propagator := otel.GetTextMapPropagator()
	parentCtx := propagator.Extract(context.TODO(), carrier)

	traceId, err := trace.TraceIDFromHex(traceIdStr)
	if err != nil {
		return nil, err
	}
	spanId, err := trace.SpanIDFromHex(spanIdStr)
	if err != nil {
		return nil, err
	}

	rsc := trace.NewSpanContext(trace.SpanContextConfig{
		TraceID:    traceId,
		SpanID:     spanId,
		TraceFlags: trace.FlagsSampled, //这个没写，是不会记录的
		TraceState: trace.TraceState{},
		Remote:     true,
	})
	// 不用pctx，不会把spanctx当做parentCtx
	sct := trace.ContextWithRemoteSpanContext(parentCtx, rsc)
	return sct, nil
}
