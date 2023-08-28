package main

import (
	"fmt"
	"github.com/richelieu-yang/chimera/v2/src/web/reqKit"
	"time"
)

func main() {
	client := reqKit.GetDefaultClient()

	client.SetTimeout(time.Second * 2)
	//ctx, _ := context.WithTimeout(context.TODO(), time.Second*2)

	resp := client.Get("http://127.0.0.1/test").Do(ctx)
	if resp.Err != nil {
		fmt.Println(reqKit.IsTimeoutError(resp.Err))
		panic(resp.Err)
	}
	fmt.Println(resp.String())
	fmt.Println(resp.ToString())
}
