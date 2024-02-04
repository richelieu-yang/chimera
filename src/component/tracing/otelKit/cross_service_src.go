package otelKit

import (
	"context"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/baggage"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/trace"
	"net/http"
)

// InjectBaggage 使用 baggage 写入 trace id 和 span id
func InjectBaggage(spanCtx context.Context, span trace.Span, r *http.Request) error {
	traceMember, err := baggage.NewMember("trace-id", span.SpanContext().TraceID().String())
	if err != nil {
		return err
	}
	spanMember, err := baggage.NewMember("span-id", span.SpanContext().SpanID().String())
	if err != nil {
		return err
	}
	b, err := baggage.New(traceMember, spanMember)
	if err != nil {
		return err
	}
	baggageCtx := baggage.ContextWithBaggage(spanCtx, b)

	//propagation.Baggage{}.Inject(baggageCtx, propagation.HeaderCarrier(r.Header))
	otel.GetTextMapPropagator().Inject(baggageCtx, propagation.HeaderCarrier(r.Header))

	return nil
}
