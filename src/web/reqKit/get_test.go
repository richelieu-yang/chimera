package reqKit

import (
	"errors"
	"fmt"
	"testing"
)

func TestGet(t *testing.T) {
	/*
		PS:
		(1) 通过判断 resp.Err 来判断是否发生 error.
		(2) 任何情况下返回的 “resp” 永远不会是 nil，可以放心的直接判断。
	*/
	//url := "https://www.baidu.com/"
	//url := "https://127.0.0.1/test"
	url := "http://127.0.0.1/测试.wps"

	client := GetDefaultClient()
	resp := client.Get(url).Do()
	if resp.Err != nil {
		panic(resp.Err)
	}
	// 不需要手动关闭
	//defer resp.Body.Close()

	if !resp.IsSuccessState() {
		panic(errors.New(fmt.Sprintf("error status: %s", resp.GetStatus())))
	}

	str := resp.String()
	fmt.Println(str)
	fmt.Println("length:", len(str))

	fmt.Println("total time:", resp.TotalTime().String())
}

func TestGet1(t *testing.T) {
	url := "https://www.baidu.com/"
	//url := "https://127.0.0.1/test"

	client := GetDefaultClient()
	resp, err := client.R().Get(url)
	if err != nil {
		panic(err)
	}
	// 不需要手动关闭
	//defer resp.Body.Close()

	if !resp.IsSuccessState() {
		panic(errors.New(fmt.Sprintf("error status: %s", resp.GetStatus())))
	}

	str := resp.String()
	fmt.Println(str)
	fmt.Println("length:", len(str))

	fmt.Println("total time:", resp.TotalTime().String())
}
