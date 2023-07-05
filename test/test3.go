package main

import (
	"context"
	"github.com/richelieu-yang/chimera/v2/src/tracing/jaegerKit"
	"github.com/sirupsen/logrus"
)

func main() {
	tp, err := jaegerKit.NewTracerProvider("http://localhost:14268/api/traces", &jaegerKit.ResourceConfig{
		ServiceName: "trace-demo",
		Environment: "production",
		ID:          1,
	})
	if err != nil {
		logrus.Fatal(err)
	}

	tracer := tp.Tracer("component-main")
	ctx, span := tracer.Start(context.TODO(), "foo")
	defer span.End()

}
