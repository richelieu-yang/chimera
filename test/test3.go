package main

import (
	"fmt"
	"go.opentelemetry.io/otel"
)

func main() {
	tp := otel.GetTracerProvider()
	fmt.Printf("%T\n", tp)

	//fmt.Println(otel.GetTracerProvider() != nil)
}
