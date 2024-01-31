package otelKit

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"net/http"
	"sync"
	"testing"
	"time"
)

func TestNewJaegerTracerProvider(t *testing.T) {
	tp, err := NewJaegerTracerProvider("http://localhost:14268/api/traces", "test", "environment", 666)
	if err != nil {
		panic(err)
	}
	otel.SetTracerProvider(tp)

	a := func(ctx context.Context, wg *sync.WaitGroup) {
		defer wg.Done()

		time.Sleep(time.Second)

		tracer := otel.Tracer("my-tracer")
		_, span := tracer.Start(ctx, "a")
		defer span.End()

		span.SetAttributes(attribute.KeyValue{
			Key:   "a",
			Value: attribute.StringValue("a"),
		})

		time.Sleep(time.Second * 2)
	}
	b := func(ctx context.Context) {
		tracer := otel.Tracer("my-tracer")
		_, span := tracer.Start(ctx, "b")
		defer span.End()

		span.SetAttributes(attribute.KeyValue{
			Key:   "b",
			Value: attribute.StringValue("b"),
		})

		time.Sleep(time.Second * 3)
	}
	engine := gin.Default()
	engine.Any("/test", func(ctx *gin.Context) {
		tracer := otel.Tracer("my-tracer")
		spanCtx, span := tracer.Start(ctx, "spanName")
		defer span.End()

		wg := new(sync.WaitGroup)
		wg.Add(1)
		go a(spanCtx, wg)
		b(spanCtx)

		wg.Wait()
		ctx.String(http.StatusOK, "hello")
	})
	if err := engine.Run(":80"); err != nil {
		logrus.Fatal(engine)
	}
}
