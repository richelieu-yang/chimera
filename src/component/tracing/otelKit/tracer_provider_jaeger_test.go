package otelKit

//import (
//	"context"
//	"github.com/gin-gonic/gin"
//	"github.com/sirupsen/logrus"
//	"go.opentelemetry.io/otel"
//	"go.opentelemetry.io/otel/attribute"
//	"net/http"
//	"sync"
//	"testing"
//	"time"
//)
//
//func TestNewJaegerTracerProvider(t *testing.T) {
//	tp, err := NewJaegerTracerProvider("http://localhost:14268/api/traces", "test", "environment", 666)
//	if err != nil {
//		panic(err)
//	}
//	otel.SetTracerProvider(tp)
//
//	funA := func(ctx context.Context, wg *sync.WaitGroup) {
//		defer wg.Done()
//
//		time.Sleep(time.Second * 2)
//
//		tracer := otel.Tracer("my-tracer")
//		_, span := tracer.Start(ctx, "funA")
//		defer span.End()
//
//		span.SetAttributes(attribute.KeyValue{
//			Key:   "funA",
//			Value: attribute.StringValue("funA"),
//		})
//
//		time.Sleep(time.Second)
//	}
//	funB := func(ctx context.Context) {
//		tracer := otel.Tracer("my-tracer")
//		_, span := tracer.Start(ctx, "funB")
//		defer span.End()
//
//		span.SetAttributes(attribute.KeyValue{
//			Key:   "funB",
//			Value: attribute.StringValue("funB"),
//		})
//
//		time.Sleep(time.Second)
//	}
//
//	engine := gin.Default()
//	engine.Any("/test", func(ctx *gin.Context) {
//		tracer := otel.Tracer("TRACER-NAME")
//		spanCtx, span := tracer.Start(ctx, "SPAN_NAME")
//		defer span.End()
//
//		wg := new(sync.WaitGroup)
//		wg.Add(1)
//		go funA(spanCtx, wg)
//		funB(spanCtx)
//		wg.Wait()
//
//		ctx.String(http.StatusOK, "hello")
//	})
//	if err := engine.Run(":80"); err != nil {
//		logrus.Fatal(engine)
//	}
//}
