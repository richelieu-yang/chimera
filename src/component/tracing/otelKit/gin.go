package otelKit

import (
	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/otel/trace"
)

// SetTracerProvider TODO:
func SetTracerProvider(ctx *gin.Context) {

}

// NewTracer TODO:
func NewTracer(ctx *gin.Context, name string, opts ...trace.TracerOption) {

}

func getParentContext(ctx *gin.Context) {

}

func setParentContext(ctx *gin.Context) {

}
