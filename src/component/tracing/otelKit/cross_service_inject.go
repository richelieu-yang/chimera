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

// InjectToHeader 使用 baggage 写入 trace id 和 span id.
/*
PS:
(1) 需要先 set up;
(2) 适用场景: 跨服务（跨应用通讯） + 发送端.
*/
func InjectToHeader(header http.Header, spanCtx context.Context, span trace.Span) error {
	carrier := propagation.HeaderCarrier(header)
	return inject(carrier, spanCtx, span)
}

// InjectToMap
/*
PS:
(1) 需要先 set up;
(2) 适用场景: 跨服务（跨应用通讯） + 发送端.
*/
func InjectToMap(m map[string]string, spanCtx context.Context, span trace.Span) error {
	carrier := propagation.MapCarrier(m)
	return inject(carrier, spanCtx, span)
}

func inject(carrier propagation.TextMapCarrier, spanCtx context.Context, span trace.Span) (err error) {
	defer func() {
		if err != nil {
			err = errorKit.Wrapf(err, "fail to inject baggage")
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

	//propagation.Baggage{}.Inject(baggageCtx, carrier)
	otel.GetTextMapPropagator().Inject(baggageCtx, carrier)

	return
}
