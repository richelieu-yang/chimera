package main

import (
	"fmt"
	rmq_client "github.com/apache/rocketmq-clients/golang"
	v2 "github.com/apache/rocketmq-clients/golang/protocol/v2"
	"github.com/richelieu42/go-scales/src/mq/rocketmq5Kit"
	"github.com/richelieu42/go-scales/src/reflectKit"
)

func main() {
	producer, err := rocketmq5Kit.NewProducer(nil, &rmq_client.Config{
		Endpoint:    "127.0.0.1:8081;127.0.0.1:8082",
		Credentials: nil,
	})
	if err != nil {
		panic(err)
	}

	v, err := reflectKit.GetNestedField(producer, "pSetting", "endpoints")
	if err != nil {
		panic(err)
	}
	endpoints := (*v2.Endpoints)(v.UnsafePointer())
	fmt.Println(endpoints.Addresses)

	//v := reflectKit.GetField(producer, "pSetting")
	//fmt.Println(v.Type())
	//v = reflect.Indirect(v)
	//v = v.FieldByName("endpoints")
	//fmt.Println(v.Type())
	//
	//fmt.Println(v.CanSet())
	//
	//ptr := (*v2.Endpoints)(v.UnsafePointer())
	//fmt.Println(ptr.Addresses)
	//
	//ptr.Addresses = []*v2.Address{
	//	&v2.Address{
	//		Host: "",
	//		Port: 1,
	//	},
	//	&v2.Address{
	//		Host: "",
	//		Port: 2,
	//	},
	//}
	//fmt.Println(ptr)
	//fmt.Println(producer)

	//v.UnsafeAddr()
	//unsafe.Pointer()

	//fmt.Println(v.Interface())

	//fmt.Println(producer)
	//
	//v := reflect.ValueOf(producer).Elem()
	//field := v.FieldByName("pSetting")
	//
	//v1 := reflect.NewAt(field.Type(), unsafe.Pointer(field.UnsafeAddr())).Elem()
	//fmt.Println(v1)
	//
	////field.Interface()
	//
	//field = v1.FieldByName("endpoints")
	//ptr := (*v2.Endpoints)(unsafe.Pointer(field.UnsafeAddr()))
	//
	//endpoints := *ptr
	//fmt.Println(endpoints)
}
