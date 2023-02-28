package main

import (
	"fmt"
	rmq_client "github.com/apache/rocketmq-clients/golang"
	v2 "github.com/apache/rocketmq-clients/golang/protocol/v2"
	"github.com/richelieu42/go-scales/src/core/strKit"
	"github.com/richelieu42/go-scales/src/mq/rocketmq5Kit"
	"reflect"
	"unsafe"
)

func main() {
	ptr := strKit.GetStringPtr("123")
	ptr = nil
	fmt.Println(strKit.GetStringPtrValue(ptr))
}

func main1() {
	producer, err := rocketmq5Kit.NewProducer(nil, &rmq_client.Config{
		Endpoint:    "127.0.0.1:8081;127.0.0.1:8082",
		Credentials: nil,
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(producer)

	v := reflect.ValueOf(producer).Elem()
	field := v.FieldByName("pSetting")

	v1 := reflect.NewAt(field.Type(), unsafe.Pointer(field.UnsafeAddr())).Elem()
	fmt.Println(v1)

	//field.Interface()

	field = v1.FieldByName("endpoints")
	ptr := (*v2.Endpoints)(unsafe.Pointer(field.UnsafeAddr()))

	endpoints := *ptr
	fmt.Println(endpoints)
}
