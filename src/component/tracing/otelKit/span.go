package otelKit

import (
	"context"
	"github.com/richelieu-yang/chimera/v3/src/core/interfaceKit"
	"github.com/richelieu-yang/chimera/v3/src/core/strKit"
	"go.opentelemetry.io/otel/trace"
)

// NewSpan
/*
@param tracer 		不能为nil
@param parentCtx 	父span的Context（没有的话，可以使用 context.TODO()）
@param spanName 	span的名称（UI页面中的Operation）
@param opts			e.g. UI页面中的Tags
						trace.WithAttributes(attribute.String("222", "222"))
@return err == nil的情况下，spanCtx 可以作为子span的传参parentCtx
*/
func NewSpan(tracer trace.Tracer, parentCtx context.Context, spanName string, opts ...trace.SpanStartOption) (spanCtx context.Context, span trace.Span, err error) {
	if err = interfaceKit.AssertNotNil(tracer, "tracer"); err != nil {
		return
	}
	if parentCtx == nil {
		parentCtx = context.TODO()
	}
	if err = strKit.AssertNotBlank(spanName, "spanName"); err != nil {
		return
	}

	spanCtx, span = tracer.Start(parentCtx, spanName, opts...)
	return
}
