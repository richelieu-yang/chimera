package otelKit

import (
	"context"
	"github.com/richelieu-yang/chimera/v2/src/core/strKit"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/sdk/trace"
)

// NewGrpcTracerProvider
/*
PS:
(1) 使用 otlptracegrpc;
(2) 此方法第二个返回值为nil的情况下，建议立即调用 otel.SetTracerProvider.

@param endpoint 	可以为""，将采用默认值: "localhost:4317"
@param serviceName 	服务名
@param attributeMap	可以为nil
@param opts 		(1) 额外配置（不要涉及 otlptracegrpc.WithEndpoint，因为在此处配了也没用）
					(2) otlptracegrpc.WithInsecure(): 	(a) 配置的话，将使用非安全协议（http...）;
														(b) 不配置的话，将使用安全协议（https...）.
					(3) 一般情况下，建议使用: otlptracegrpc.WithInsecure(), otlptracegrpc.WithDialOption(grpc.WithBlock())
*/
func NewGrpcTracerProvider(endpoint, serviceName string, attributeMap map[string]string, opts ...otlptracegrpc.Option) (*trace.TracerProvider, error) {
	// 放在最后面（优先级最高）
	if strKit.IsNotEmpty(endpoint) {
		opts = append(opts, otlptracegrpc.WithEndpoint(endpoint))
	}
	// 创建 exporter 实例
	exporter, err := otlptracegrpc.New(context.TODO(), opts...)
	if err != nil {
		return nil, err
	}

	res, err := newDetailedResource(serviceName, attributeMap)
	if err != nil {
		return nil, err
	}

	tp := trace.NewTracerProvider(
		/*
			Always be sure to batch in production.

			trace.WithBatcher(): 	是一个配置选项，用于设置 Span 的导出方式。
									当使用 trace.WithBatcher() 时，Span 是通过通道异步按批发送的。
									默认情况下，每隔一段时间（默认批处理5s超时触发）发送，或者当达到最大批处理长度（默认512）时，会把要批处理的 Span 发送出去。
		*/
		trace.WithBatcher(exporter),

		/*
			Record information about this application in a Resource.

			trace.WithResource(): 	是一个配置选项，用于设置 TracerProvider 的资源信息。
									资源是一组描述 TracerProvider 的属性，例如服务名称、环境、地理位置等。这些属性可以帮助你更好地理解和过滤跟踪数据1。
		*/
		trace.WithResource(res),

		/*
			trace.WithSampler():	是一个配置选项，用于设置采样器的行为，采样器决定了哪些跟踪信息应该被记录和导出.
			trace.AlwaysSample():	创建了一个采样器，它会选择所有的跟踪进行采样。
									这意味着每一个创建的 Span 都会被记录和可能被导出。这对于调试非常有用，因为你可以看到所有的跟踪信息。
									然而，在生产环境中，这可能会产生大量的数据，因此可能需要使用更复杂的采样策略。
		*/
		trace.WithSampler(trace.AlwaysSample()),
		//trace.WithSampler(trace.NeverSample()),
	)
	return tp, nil
}
