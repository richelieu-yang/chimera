package otelKit

import (
	"context"
	"github.com/richelieu-yang/chimera/v3/src/component/web/httpKit"
	"github.com/richelieu-yang/chimera/v3/src/core/errorKit"
	"github.com/richelieu-yang/chimera/v3/src/core/strKit"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/baggage"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/trace"
	"net/http"
)

// GetRemoteSpanCtx
/*
PS:
(1) 适用场景: 跨服务 + 接收端
(2) 返回的error != nil的情况下，可以判断是否等于 otelKit.NotOtelRequestError.
*/
func GetRemoteSpanCtx(r *http.Request) (remoteSpanCtx context.Context, err error) {
	baggageStr := httpKit.GetHeader(r.Header, HeaderBaggage)
	if strKit.IsEmpty(baggageStr) {
		// 非链路追踪请求
		err = NotOtelRequestError
		return
	}

	defer func() {
		if err != nil {
			err = errorKit.Wrap(err, "Fail to get remote span context")
		}
	}()

	//var propagator = propagation.TextMapPropagator(propagation.Baggage{})
	//propagatorCtx := propagator.Extract(r.Context(), propagation.HeaderCarrier(r.Header))
	propagatorCtx := otel.GetTextMapPropagator().Extract(r.Context(), propagation.HeaderCarrier(r.Header))

	b := baggage.FromContext(propagatorCtx)
	traceId, err := trace.TraceIDFromHex(b.Member(KeyTraceId).Value())
	if err != nil {
		return
	}
	spanId, err := trace.SpanIDFromHex(b.Member(KeySpanId).Value())
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
