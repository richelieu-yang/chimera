package otelKit

import (
	"context"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"
)

type (
	Tracer struct {
		origin trace.Tracer
	}
)

func (tracer *Tracer) NewSpanAndStart(parentSpanCtx context.Context, spanName string, opts ...trace.SpanStartOption) (context.Context, trace.Span) {
	if tracer.origin == nil {

	}
	tracer.origin.Start(parentSpanCtx, spanName, opts...)

}

// NewTracer
/*
@param tracerName 建议为""
*/
func NewTracer(tracerName string, opts ...trace.TracerOption) *Tracer {
	var origin trace.Tracer

	if err := check(); err != nil {
		// 无效的Tracer
		origin = nil
	} else {
		origin = otel.Tracer(tracerName, opts...)
	}
	return &Tracer{
		origin: origin,
	}
}
