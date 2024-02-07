package main

func main() {
	///* (1) 初始获取到的值不为nil */
	//tp := otel.GetTracerProvider()
	//fmt.Println(tp)        // &{<nil> {0 0} map[] <nil>}
	//fmt.Println(tp != nil) // true
	//
	///* (2) 设置为nil后，再次获取得到nil */
	//var tp1 trace.TracerProvider = nil
	//otel.SetTracerProvider(tp1)
	//fmt.Println(otel.GetTracerProvider() == nil) // true
}
