package main

import (
	"fmt"
	"go.opentelemetry.io/otel"
)

func main() {
	tp := otel.GetTracerProvider()

	fmt.Println(tp != nil)        // true
	fmt.Printf("%T %v\n", tp, tp) // *global.tracerProvider &{<nil> {0 0} map[] <nil>}
}
