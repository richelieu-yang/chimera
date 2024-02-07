package otelKit

import (
	"context"
	"github.com/richelieu-yang/chimera/v3/src/core/errorKit"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/baggage"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/trace"
	"net/http"
)

// InjectBaggage 使用 baggage 写入 trace id 和 span id
/*
适用场景: 跨服务 + 发送端
*/
func InjectBaggage(r *http.Request, spanCtx context.Context, span trace.Span) (err error) {
	defer func() {
		if err != nil {
			err = errorKit.Wrap(err, "Fail to inject baggage")
		}
	}()

	traceMember, err := baggage.NewMember(KeyTraceId, span.SpanContext().TraceID().String())
	if err != nil {
		return
	}
	spanMember, err := baggage.NewMember(KeySpanId, span.SpanContext().SpanID().String())
	if err != nil {
		return
	}
	b, err := baggage.New(traceMember, spanMember)
	if err != nil {
		return
	}
	baggageCtx := baggage.ContextWithBaggage(spanCtx, b)

	//propagation.Baggage{}.Inject(baggageCtx, propagation.HeaderCarrier(r.Header))
	otel.GetTextMapPropagator().Inject(baggageCtx, propagation.HeaderCarrier(r.Header))

	return
}
