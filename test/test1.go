package main

import (
	"fmt"
	"github.com/richelieu42/chimera/src/funcKit"
)

func main() {
	testFunc()
	//addresses := []string{"localhost:6650"}
	//
	//ctx, cancel := context.WithTimeout(context.TODO(), time.Second*10)
	//defer cancel()
	//consumer, err := pulsarKit.NewConsumerOriginally(ctx, addresses, pulsar.ConsumerOptions{
	//	Topic:            "test",
	//	SubscriptionName: "my-sub1",
	//	Type:             pulsar.Exclusive,
	//}, "")
	//if err != nil {
	//	logrus.Fatal(err)
	//}
	//logrus.Info(consumer)
}

func testFunc() {
	fmt.Println(funcKit.GetEntireCaller(1))
	fmt.Println(funcKit.GetCaller(1))
}
