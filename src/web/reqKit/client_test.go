package reqKit

import (
	"errors"
	"fmt"
	"testing"
)

func TestNewClient(t *testing.T) {
	{
		/*

			PS:
			(1) 通过判断 resp.Err 来判断是否发生 error.
			(2) 任何情况下返回的 “resp” 永远不会是 nil，可以放心的直接判断。
		*/
		url := "https://www.baidu.com/"

		client := NewClient()
		resp := client.Get(url).Do()
		if resp.Err != nil {
			panic(resp.Err)
		}
		defer func() {
			if resp.Body != nil {
				_ = resp.Body.Close()
			}
		}()

		if !resp.IsSuccessState() {
			panic(errors.New(fmt.Sprintf("error status: %s", resp.GetStatus())))
		}

		str := resp.String()
		fmt.Println("length:", len(str))

		fmt.Println("total time:", resp.TotalTime().String())
	}

}
