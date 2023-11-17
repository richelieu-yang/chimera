package gaodeKit

import (
	"fmt"
	"github.com/richelieu-yang/chimera/v2/src/json/jsonKit"
	"testing"
)

func TestClient_GetIp(t *testing.T) {
	client, err := NewClient("b15c36bf1df4c272e92f3f1875a127f1")
	if err != nil {
		panic(err)
	}

	//ip := "218.90.174.146"
	ip := "10.0.9.141"
	ipInfo, err := client.GetIp(ip)
	if err != nil {
		panic(err)
	}
	fmt.Println(jsonKit.MarshalIndentToString(ipInfo, "", "    "))
}
