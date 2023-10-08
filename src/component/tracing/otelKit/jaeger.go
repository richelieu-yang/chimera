package otelKit

import (
	"context"
	"github.com/sirupsen/logrus"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/jaeger"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.17.0"
	"time"
)

// NewJaegerTracerProvider
/*
PS:
(1) agent		默认端口: 6831
(2) collector	默认端口: 14268
(3) 创建 TracerProvider 成功的情况下（即第2个返回值为nil），建议调用 otel.SetTracerProvider().

@param url 连接的jaeger服务（agent || collector） e.g."http://localhost:14268/api/traces"
*/
func NewJaegerTracerProvider(url, service, environment string, id int64) (*trace.TracerProvider, error) {
	/*
		创建jaeger provider.
		PS: 可以直接连collector，也可以连agent
	*/
	exporter, err := jaeger.New(jaeger.WithCollectorEndpoint(jaeger.WithEndpoint(url)))
	if err != nil {
		return nil, err
	}
	tp := trace.NewTracerProvider(
		// Always be sure to batch in production.
		trace.WithBatcher(exporter),
		// Record information about this application in a Resource.
		trace.WithResource(resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceNameKey.String(service),
			attribute.String("environment", environment),
			attribute.Int64("ID", id),
		)),
	)
	return tp, nil
}

func ShutdownJaegerTracerProvider(tp *trace.TracerProvider) {
	if tp == nil {
		return
	}

	ctx, cancel := context.WithTimeout(context.TODO(), time.Second*5)
	defer cancel()
	if err := tp.Shutdown(ctx); err != nil {
		logrus.WithError(err).Error("Fail to shutdown Jaeger tracer provider.")
	}
}
