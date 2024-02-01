package main

import (
	"fmt"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
)

func main() {
	s := []int{0, 1, 2, 3, 4, 5}
	a := [3]int(s[:])

	fmt.Printf("%T %v\n", s, s) // []int [0 1 2 3 4 5]
	fmt.Printf("%T %v\n", a, a) // [3]int [0 1 2]

	otlptracehttp.NewClient(otlptracehttp.WithEndpoint("localhost:4318"))

}
